package network

import (
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
	tcpListener     net.Listener
	OnSessionPacket func(packet *SessionPacket)
}

func NewServer(address string) *Server {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcp", resolveTCPAddr)
	if err != nil {
		panic(err)
	}
	s := &Server{}
	s.tcpListener = tcpListener
	return s
}

func (s *Server) Run() {
	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				continue
			}
		}
		go func() {
			newSession := NewSession(conn)
			SessionMgrInstance.AddSession(newSession)
			newSession.Run()
			SessionMgrInstance.DelSession(newSession.UID)
		}()
	}
}
