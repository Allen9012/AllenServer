package world

import (
	"github.com/Allen9012/AllenServer/aop/logger"
	"github.com/Allen9012/AllenServer/business/module/chat"
	"github.com/Allen9012/AllenServer/business/module/player"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"os"
	"syscall"
)

type World struct {
	Pm              *player.Manager
	Server          *network.Server
	Handlers        map[messageID.MessageId]func(message *network.Packet)
	chSessionPacket chan *network.Packet
	ChatSystem      chat.System
}

func NewWorld() *World {
	m := &World{Pm: player.NewPlayerMgr()}
	m.Server = network.NewServer(":8023", 100, 200, logger.GetLogger())
	m.Server.MessageHandler = m.OnSessionPacket
	m.Handlers = make(map[messageID.MessageId]func(message *network.Packet))
	m.ChatSystem.SetOwner(m)
	return m
}

var Oasis *World

func (w *World) Start() {
	w.HandlerRegister()
	go w.Server.Run()
	go w.Pm.Run()
}

func (w *World) Stop() {

}

func (w *World) OnSessionPacket(packet *network.Packet) {
	if handler, ok := w.Handlers[messageID.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
		return
	}
	if p := w.Pm.GetPlayer(uint64(packet.Conn.ConnID)); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

func (w *World) OnSystemSignal(signal os.Signal) bool {
	logger.Debug("[World] 收到信号 %v \n", signal)
	tag := true
	switch signal {
	case syscall.SIGHUP:
		//todo
	case syscall.SIGPIPE:
	default:
		logger.Debug("[World] 收到信号准备退出...")
		tag = false

	}
	return tag
}
