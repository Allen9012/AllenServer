package gateway

import (
	"github.com/Allen9012/AllenServer/network"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc: 主要为了维护客户端的链接
  @modified by:
**/

type ClientServer struct {
	real    *network.Server
	session sync.Map //*network.Session
	//	与其他服务器绑定信息
	FromInnerCh chan interface{}
	ToInnerCh   chan interface{}
}

func NewClientServer() *ClientServer {
	return &ClientServer{
		real:    network.NewServer(""),
		session: sync.Map{},
	}
}

func (c *ClientServer) loop() {
	c.real.OnSessionPacket = c.MessageHandler
	for {
		select {
		case data := <-c.FromInnerCh:
			c.Router(data)
		}
	}
}

func (s *ClientServer) MessageHandler(packet *network.SessionPacket) {
	//todo check
	s.ToInnerCh <- packet
}

// Router 路由给服务器
func (s *ClientServer) Router(data interface{}) {
	//handler := s.router.GetHandler(data.(*network.Packet))
	//if handler != nil {
	//	val := data.(*network.Packet)
	//	handler(val, nil) //todo
	//}
}
