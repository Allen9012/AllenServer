package gateway

import (
	"github.com/Allen9012/AllenServer/gateway/fuse"
	"github.com/Allen9012/AllenServer/network"
	"sync"
	"sync/atomic"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc: 主要为了维护客户端的链接
  @modified by:
**/

var (
	ClientServerInstance *ClientServer
	onceInit             sync.Once
)

type ClientServer struct {
	real        *network.Server
	FromInnerCh chan interface{} //与其他服绑定信息
	ToInnerCh   chan interface{}
	router      *fuse.Router
	clients     sync.Map
}

type ClientInfo struct {
	onlineID atomic.Value
	userId   uint64
	conn     *network.TcpConnX
}

func GetClientServerInstance() *ClientServer {
	onceInit.Do(func() {
		ClientServerInstance = &ClientServer{}
	})
	return ClientServerInstance
}

func NewClientServer() *ClientServer {
	c := &ClientServer{
		//real:     network.NewServer(""),
	}
	c.router = fuse.NewRouter()
	c.router.Use(CheckPacketSecurity) //中间件
	return c
}

func (s *ClientServer) loop() {
	s.real.MessageHandler = s.MessageHandler
	for {
		select {
		case data := <-s.FromInnerCh:
			s.Router(data)
		}
	}

}
func (s *ClientServer) MessageHandler(packet *network.Packet) {
	//todo check
	s.ToInnerCh <- packet
}

// Router 路由给服务器
func (s *ClientServer) Router(data interface{}) {
	handler := s.router.GetHandler(data.(*network.Packet))
	if handler != nil {
		val := data.(*network.Packet)
		handler(val, nil) //todo
	}
}

func (s *ClientServer) Register() {
	s.router.AddRoute(333, s.RegisterLoginInfo)
	s.router.AddRoute(444, s.ForwardServerPacket)
}

func (s *ClientServer) bindUserId2Client(userId uint64, conn *network.TcpConnX) {
	_, ok := s.clients.Load(userId)
	if ok {
		//todo
	}
	c := &ClientInfo{
		onlineID: atomic.Value{},
		userId:   userId,
		conn:     conn,
	}
	s.clients.Store(userId, c)
}

func (s *ClientServer) unbindUserId2Client(userId uint64, conn *network.TcpConnX) {
	ci, ok := s.clients.Load(userId)
	if !ok {
		return
	}

	if ci.(*ClientInfo).conn == conn {
		s.clients.Delete(userId)
	}
}

// id创建客户端连接
func (s *ClientServer) getClientWithUserId(userId uint64) *ClientInfo {
	client, ok := s.clients.Load(userId)
	if !ok {
		return nil
	}
	return client.(*ClientInfo)
}

// 断开连接
func (s *ClientServer) onlineServerDisconnected(srvID, srvAddr string, zoneId int, proIndx uint32) {
	s.clients.Range(func(k, v interface{}) bool {
		ci, ok := v.(*ClientInfo)
		if ok && ci.onlineID.Load().(string) == srvID {
			if ci.userId != 0 {
			}
			s.unbindUserId2Client(ci.userId, ci.conn)
			ci.userId = 0
			ci.conn.Close()
		}
		return true
	})
}
