package network

import (
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
	UID            uint64
	Conn           net.Conn
	IsClose        bool
	packer         IPacker
	WriteCh        chan *Message
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
	//
}

// NewSession 网络通信默认是大端
func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn:    conn,
		packer:  NormalPackerInstance,
		WriteCh: make(chan *Message, 1),
	}
}

// Run 处理收发数据
func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Read() {
	// 一直读取数据
	for {
		err := s.Conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			continue
		}
		message, err := s.packer.Unpack(s.Conn)
		if _, ok := err.(net.Error); ok {
			continue
		}
		fmt.Println("server receive message:", string(message.Data))
		// 处理消息
		s.MessageHandler(&SessionPacket{
			Msg:  message,
			Sess: s,
		})
	}
}

// Write 持续将WriteCh信息发送出去
func (s *Session) Write() {
	for {
		select {
		case resp := <-s.WriteCh:
			s.send(resp)
		}
	}
}

// 超时时间需要加在send中，而不是Write中
func (s *Session) send(message *Message) {
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := s.packer.Pack(message)
	if err != nil {
		return
	}

	_, err = s.Conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Session) SendMsg(msg *Message) {
	s.WriteCh <- msg
}