package network

import "github.com/Allen9012/AllenServer/network/example"

type ISession interface {
	OnConnect()
	OnClose()
	OnMessage(*Message, *example.TcpSession)
}
