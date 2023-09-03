package gateway

import (
	"github.com/Allen9012/AllenServer/gateway/fuse"
	"github.com/Allen9012/AllenServer/network"
)

func (s *InnerServer) ForwardClientPacket(packet *network.Packet, principal fuse.Principal) {
	// 转发消息给客户端
}

func (s *InnerServer) ServerInfoRegister(packet *network.Packet, principal fuse.Principal) {
	// 服务器信息的注册
}
