package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:  模拟客户端收发请求
	@modified by:
*/

type Client struct {
	Address   string
	packer    IPacker
	ChMsg     chan *Message
	OnMessage func(message *ClientPacket)
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
		packer: &NormalPacker{
			Order: binary.BigEndian,
		},
		ChMsg: make(chan *Message, 1),
	}
}
func (c *Client) Run() {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	客户端可以监听多个conn，所以不放在结构体中
	go c.Write(conn)
	go c.Read(conn)
}

func (c *Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			c.ChMsg <- &Message{
				ID:   111,
				Data: []byte("client send msg"),
			}
		case msg := <-c.ChMsg:
			c.Send(conn, msg)
		}
	}
}

func (c *Client) Send(conn net.Conn, message *Message) {
	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.Unpack(conn)
		if _, ok := err.(net.Error); ok && err != nil {
			fmt.Println(err)
			continue
		}
		c.OnMessage(&ClientPacket{
			Msg:  message,
			Conn: conn,
		})
		fmt.Println("resp message:", string(message.Data))
	}
}
