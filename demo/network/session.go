package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 维护连接
  @modified by:
**/

// Session 连接
type Session struct {
	conn    net.Conn
	packer  *NormalPacker
	chWrite chan *Message
}

// NewSession 网络通信默认是大端
func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		chWrite: make(chan *Message, 1),
	}
}

// Run 处理收发数据
func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Close() {

}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	// 一直读取数据
	for {
		message, err := s.packer.Unpack(s.conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("server receive message:", string(message.Data))
		s.chWrite <- &Message{
			ID:   999,
			Data: []byte("receive message"),
		}
	}
}

func (s *Session) Write() {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case msg := <-s.chWrite:
			s.send(msg)
		}
	}
}

func (s *Session) send(message *Message) {
	bytes, err := s.packer.Pack(message)
	if err != nil {
		return
	}

	_, err = s.conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
