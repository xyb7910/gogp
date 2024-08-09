package net

import (
	"encoding/binary"
	"net"
	"testing"
)

// 最简单的服务端
func TestServerV1(t *testing.T) {
	// 开始监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	// 开始接收连接
	conn, err := listener.Accept()
	if err != nil {
		t.Fatal(err)
	}
	handleV1(conn)

}

func handleV1(conn net.Conn) {
	_, err := conn.Write([]byte("hello world"))
	if err != nil {
		conn.Close()
		return
	}
}

func TestServerV2(t *testing.T) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			handleV2(conn)
		}()
	}
}

func handleV2(conn net.Conn) {
	for {
		lenBs := make([]byte, 8)
		_, err := conn.Read(lenBs)
		if err != nil {
			conn.Close()
			return
		}
		msgLen := binary.BigEndian.Uint64(lenBs)
		reqBs := make([]byte, msgLen)
		_, err = conn.Read(reqBs)
		if err != nil {
			conn.Close()
			return
		}
		_, err = conn.Write([]byte("hello world"))
		if err != nil {
			conn.Close()
			return
		}
	}
}
