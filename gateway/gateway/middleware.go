package gateway

import (
	"github.com/Allen9012/AllenServer/gateway/fuse"
	"github.com/Allen9012/AllenServer/network"
)

func CheckPacketSecurity(handler fuse.Handler) fuse.Handler {
	return func(packet *network.Packet, p fuse.Principal) {
		//todo check packet security
		handler(packet, p)
	}
}

// TODO流量限制
