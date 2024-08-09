package net

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// 学习： https://blog.csdn.net/qq_39383767/article/details/130027692

type Option func(p *SimplePool)

type conn struct {
	c          net.Conn
	lastActive time.Time
}

type conReq struct {
	con chan conn
}

type SimplePool struct {
	idleChan chan conn    // 空闲连接
	waitChan chan *conReq // 阻塞的请求

	factory     func() (net.Conn, error) // 连接工厂
	idleTimeOut time.Duration            // 空闲时间

	maxCnt int32 // 最大连接数
	cnt    int32 // 当前连接数

	l sync.Mutex // 锁
}

func NewSimplePool(factory func() (net.Conn, error), opt ...Option) *SimplePool {
	res := &SimplePool{
		idleChan: make(chan conn, 100),
		waitChan: make(chan *conReq, 100),
		factory:  factory,
		maxCnt:   100,
	}
	for _, o := range opt {
		o(res)
	}
	return res
}

// WithMaxIdleCnt 设置最大空闲连接数
func WithMaxIdleCnt(maxIdleCnt int32) Option {
	return func(p *SimplePool) {
		p.idleChan = make(chan conn, maxIdleCnt)
	}
}

// WithMaxCnt 设置最大连接数量
func WithMaxCnt(maxCnt int32) Option {
	return func(p *SimplePool) {
		p.maxCnt = maxCnt
	}
}

func (p *SimplePool) Get() (net.Conn, error) {
	for {
		select {
		// 当有空闲的时候，首先从空闲队列中取
		case c := <-p.idleChan:
			// 如果空闲时间超过了，则关闭连接
			if c.lastActive.Add(p.idleTimeOut).Before(time.Now()) {
				// 使用原子操作减少计数
				atomic.AddInt32(&p.cnt, -1)
				_ = c.c.Close()
				continue
			}
			return c.c, nil
		// 如果没有空闲的，则创建新的连接
		default:
			// 判断是否超过最大连接数
			cnt := atomic.AddInt32(&p.cnt, 1)
			// 如果没有超过，则创建新的连接
			if cnt <= p.maxCnt {
				return p.factory()
			}
			// 如果超过了，则等待
			atomic.AddInt32(&p.cnt, -1)
			req := &conReq{
				con: make(chan conn, 1),
			}
			p.waitChan <- req
			c := <-req.con
			return c.c, nil
		}
	}
}

func (p *SimplePool) Put(c net.Conn) {
	// 如果有阻塞的，直接转交连接
	p.l.Lock()
	if len(p.waitChan) > 0 {
		req := <-p.waitChan
		p.l.Unlock()
		req.con <- conn{
			c:          c,
			lastActive: time.Now(),
		}
		return
	}
	p.l.Unlock()
	// 没有阻塞时候
	select {
	// 如果空闲队列没有满，则放入空闲队列
	case p.idleChan <- conn{c: c, lastActive: time.Now()}:
	default:
		// 如果满了，则关闭连接
		defer func() {
			atomic.AddInt32(&p.cnt, -1)
		}()
		_ = c.Close()
	}
}
