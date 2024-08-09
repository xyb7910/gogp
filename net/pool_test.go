package net

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

// mockConn 用于辅助测试
type mockConn struct {
	closed bool
}

func TestSimplePoolPool(t *testing.T) {
	p := NewSimplePool(func() (net.Conn, error) {
		return &mockConn{}, nil
	}, WithMaxIdleCnt(2), WithMaxCnt(3))

	// 这三个连接都可以正常获取
	c1, err := p.Get()
	assert.Nil(t, err)
	c2, err := p.Get()
	assert.Nil(t, err)
	c3, err := p.Get()
	assert.Nil(t, err)

	// 放回两个连接
	p.Put(c1)
	p.Put(c2)

	// 空闲队列满了，关闭连接
	p.Put(c3)

	assert.True(t, c3.(*mockConn).closed)
}

func (m mockConn) Read(b []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mockConn) Write(b []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mockConn) Close() error {
	fmt.Println("close conn")
	m.closed = true
	return nil
}

func (m mockConn) LocalAddr() net.Addr {
	//TODO implement me
	panic("implement me")
}

func (m mockConn) RemoteAddr() net.Addr {
	//TODO implement me
	panic("implement me")
}

func (m mockConn) SetDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (m mockConn) SetReadDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (m mockConn) SetWriteDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}
