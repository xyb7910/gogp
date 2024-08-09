package net

import (
	"encoding/binary"
	"net"
	"testing"
)

// 最简单的客户端
func TestClientV1(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		conn.Close()
		return
	}

}

// 实现一个可以发送大数据的客户端
func TestClientV2(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	// 表示消息
	msg := "How are you?"
	msgLen := len(msg)
	msgLenBs := make([]byte, 8)
	binary.BigEndian.PutUint64(msgLenBs, uint64(msgLen))
	data := append(msgLenBs, []byte(msg)...)
	_, err = conn.Write(data)
	if err != nil {
		conn.Close()
		return
	}

	respBs := make([]byte, 16)
	_, err = conn.Read(respBs)
	if err != nil {
		conn.Close()
	}
}
