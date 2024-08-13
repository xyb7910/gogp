package net_pro

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func Connect(network, address string) error {
	conn, err := net.DialTimeout(network, address, 3*time.Second)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	for i := 0; i < 10; i++ {
		_, err = conn.Write([]byte("hello world"))
		if err != nil {
			return nil
		}
		res := make([]byte, 128)
		_, err = conn.Read(res)
		if err != nil {
			return nil
		}
		fmt.Println(string(res))
	}
	return nil
}

type Client struct {
	network string
	address string
}

func (c *Client) Send(data string) (string, error) {
	conn, err := net.DialTimeout(c.network, c.address, 3*time.Second)
	if err != nil {
		return "", nil
	}
	defer func() {
		_ = conn.Close()
	}()

	reqLen := len(data)

	// data = reqLen 的64位表示 + respData
	req := make([]byte, reqLen+numOfLengthBytes)
	// 先把长度写入到req中
	binary.BigEndian.PutUint64(req[:numOfLengthBytes], uint64(reqLen))
	// 再把数据写入到req中
	copy(req[numOfLengthBytes:], data)

	_, err = conn.Write(req)
	if err != nil {
		return "", err
	}

	lenBs := make([]byte, numOfLengthBytes)
	_, err = conn.Read(lenBs)
	if err != nil {
		return "", nil
	}
	// 响应的长度
	length := binary.BigEndian.Uint64(lenBs)

	respBs := make([]byte, length)
	_, err = conn.Read(respBs)
	if err != nil {
		return "", nil
	}
	return string(respBs), nil
}
