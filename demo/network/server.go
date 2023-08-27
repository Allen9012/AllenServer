package network

import (
	"fmt"
	"net"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 处理客户端连接
  @modified by:
**/

type Server struct {
	listener net.Listener
	address  string
	network  string
}

func NewServer(address, network string) *Server {
	return &Server{
		listener: nil,
		address:  address,
		network:  network,
	}
}

// TODO 日志模块

// Run 运行服务器
func (s *Server) Run() {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.address)
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.listener = tcpListener
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		go func() {
			new_session := NewSession(conn)
			new_session.Run()
		}()
	}
}
