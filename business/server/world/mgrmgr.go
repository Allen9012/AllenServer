package world

import (
	"github.com/Allen9012/AllenServer/aop/logger"
	"github.com/Allen9012/AllenServer/business/module/player"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"os"
	"syscall"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 保存所有的manager
  @modified by:
**/

type MgrMgr struct {
	Pm              *player.Manager
	Server          *network.Server
	Handlers        map[messageID.MessageId]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

var MM *MgrMgr

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{Pm: player.NewPlayerMgr()}
	m.Server = network.NewServer(":8888")
	m.Server.OnSessionPacket = m.OnSessionPacket
	m.Handlers = make(map[messageID.MessageId]func(message *network.SessionPacket))
	return m
}

func (mm *MgrMgr) Start() {
	mm.HandlerRegister()
	go mm.Server.Run()
	go mm.Pm.Run()
}

func (mm *MgrMgr) Stop() {

}

// OnSessionPacket 判读消息如果注册了，则可以handle
func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := mm.Handlers[messageID.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
		return
	}
	// 发送给player
	if p := mm.Pm.GetPlayer(packet.Sess.UID); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

func (mm *MgrMgr) OnSystemSignal(signal os.Signal) bool {
	logger.Debug("[MgrMgr] 收到信号 %v \n", signal)
	tag := true
	switch signal {
	case syscall.SIGHUP:
		//todo
	case syscall.SIGPIPE:
	default:
		logger.Debug("[MgrMgr] 收到信号准备退出...")
		tag = false

	}
	return tag
}
