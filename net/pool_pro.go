package net

import (
	"errors"
	"golang.org/x/net/context"
	"net"
	"sync"
	"time"
)

type idleConn struct {
	c              net.Conn
	lastActiveTime time.Time
}

type connReq struct {
	connChan chan net.Conn
}

type Pool struct {
	// 空闲连接队列
	idlesConns chan *idleConn
	// 请求队列
	reqQueue []connReq
	// 最大连接数
	maxCnt int
	// 当前连接数
	cnt int
	// 连接超时时间
	maxIdleTime time.Duration

	factory func() (net.Conn, error)

	lock sync.Mutex
}

func NewPool(initCnt int, maxIdleCnt int, maxCnt int,
	maxIdleTime time.Duration,
	factory func() (net.Conn, error)) (*Pool, error) {
	if initCnt > maxIdleCnt {
		return nil, errors.New("initCnt > maxIdleCnt")
	}
	idlesConns := make(chan *idleConn, maxIdleCnt)
	for i := 0; i < initCnt; i++ {
		conn, err := factory()
		if err != nil {
			return nil, err
		}
		idlesConns <- &idleConn{
			c:              conn,
			lastActiveTime: time.Now(),
		}
	}
	res := &Pool{
		idlesConns:  idlesConns,
		maxCnt:      maxCnt,
		cnt:         0,
		maxIdleTime: maxIdleTime,
		factory:     factory,
	}
	return res, nil
}

func (p *Pool) Get(ctx context.Context) (net.Conn, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	for {
		select {
		case ic := <-p.idlesConns:
			// 拿到了空闲连接
			if ic.lastActiveTime.Add(p.maxIdleTime).Before(time.Now()) {
				_ = ic.c.Close()
				continue
			}
			return ic.c, nil
		default:
			// 没有空闲连接
			p.lock.Lock()
			if p.cnt >= p.maxCnt {
				// 连接数已经达到上限, 等待
				req := connReq{connChan: make(chan net.Conn, 1)}
				p.reqQueue = append(p.reqQueue, req)
				p.lock.Unlock()
				select {
				case <-ctx.Done():
					go func() {
						c := <-req.connChan
						_ = p.Put(c)
					}()
					return nil, ctx.Err()
				case c := <-req.connChan:
					return c, nil
				}
			}
			c, err := p.factory()
			if err != nil {
				return nil, err
			}
			p.cnt++
			p.lock.Unlock()
			return c, nil
		}
	}
}

func (p *Pool) Put(c net.Conn) error {
	p.lock.Lock()
	if len(p.reqQueue) > 0 {
		req := p.reqQueue[0]
		p.reqQueue = p.reqQueue[1:]
		p.lock.Unlock()
		req.connChan <- c
		return nil
	}
	defer p.lock.Unlock()

	ic := &idleConn{
		c:              c,
		lastActiveTime: time.Now(),
	}
	select {
	case p.idlesConns <- ic:
	default:
		_ = c.Close()
		p.cnt--
	}
	return nil
}
