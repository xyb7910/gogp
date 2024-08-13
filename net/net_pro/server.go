package net_pro

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const numOfLengthBytes = 8

func Serve(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		fmt.Errorf("net.Listen error:%v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("listener.Accept error:%v", err)
		}
		go func() {
			if er := handleConnection(conn); er != nil {
				_ = conn.Close()
			}
		}()
	}
}

func handleConnection(conn net.Conn) error {
	for {
		bs := make([]byte, 8)
		_, err := conn.Read(bs)
		if err == net.ErrClosed || err == io.EOF || err == io.ErrUnexpectedEOF {
			return err
		}
		if err != nil {
			continue
		}
		res := handleMsg(bs)
		_, err = conn.Write(res)
		if err == net.ErrClosed || err == io.EOF || err == io.ErrUnexpectedEOF {
			return err
		}
		if err != nil {
			continue
		}
	}

}

func handleMsg(req []byte) []byte {
	res := make([]byte, 2*len(req))
	copy(res[:len(req)], req)
	copy(res[len(req):], req)
	return res
}

type Server struct {
	network string
	address string
}

func NewServer(network, address string) *Server {
	return &Server{
		network: network,
		address: address,
	}
}

func (s *Server) Start(network, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			if er := s.handleConn(conn); er != nil {
				_ = conn.Close()
			}
		}()
	}
}

func (s *Server) handleConn(conn net.Conn) error {
	for {
		lenBs := make([]byte, numOfLengthBytes)
		_, err := conn.Read(lenBs)
		if err != nil {
			return err
		}

		// 读取消息长度
		length := binary.BigEndian.Uint64(lenBs)

		reqBs := make([]byte, length)
		_, err = conn.Read(reqBs)
		if err != nil {
			return err
		}
		respData := handleMsg(reqBs)
		respLen := len(respData)

		// 构建响应数据
		// data = 消息长度 + 消息内容
		res := make([]byte, numOfLengthBytes+respLen)
		// 写入消息长度
		binary.BigEndian.PutUint64(res, uint64(respLen))
		// 写入消息内容
		copy(res[numOfLengthBytes:], respData)
		_, err = conn.Write(res)
		if err != nil {
			return err
		}
	}
}
