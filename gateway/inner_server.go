package gateway

import (
	"github.com/Allen9012/AllenServer/network"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type InnerServer struct {
	real     *network.Server
	sessions sync.Map //*network.Session
	// 哪些消息交给这些服务器处理（messageID控制范围）
	FromClientCh chan interface{}
	ToClientCh   chan interface{}
}

func NewInnerServer() *InnerServer {
	return &InnerServer{
		real: network.NewServer(""),
	}
}

func (s *InnerServer) loop() {
	s.real.OnSessionPacket = s.MessageHandler
	s.real.Run()
	for {
		select {
		case data := <-s.FromClientCh:
			s.Router(data)
		}
	}
}

func (s *InnerServer) MessageHandler(packet *network.SessionPacket) {
	//如果是注册节点信息

	s.ToClientCh <- packet
}

func (s *InnerServer) AddSession(session *network.Session) {
	s.sessions.Store(session, session)
}

func (s *InnerServer) DeleteSession(session *network.Session) {
	s.sessions.Delete(session)
}

// Router 路由到指定客户端处理
func (s *InnerServer) Router(data interface{}) {
	////world server 多节点支持
	////一般逻辑都会经过world server 做转发
	////战斗服的话，会直连客户端
	//
	////get serverUId
	//handler := s.router.GetHandler(data.(*network.Packet))
	//if handler != nil {
	//	val := data.(*network.Packet)
	//	handler(val, nil)
	//}
	////todo 发送给对应的服务器处理
}
