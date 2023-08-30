package world

import (
	"github.com/Allen9012/AllenServer/demo/manager"
	"github.com/Allen9012/AllenServer/demo/network"
	"github.com/Allen9012/AllenServer/demo/network/protocol/gen/messageID"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 保存所有的manager
  @modified by:
**/

type MgrMgr struct {
	Pm              *manager.PlayerMgr
	Server          *network.Server
	Handlers        map[messageID.MessageId]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

var MM *MgrMgr

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{Pm: &manager.PlayerMgr{}}
	m.Server = network.NewServer(":8888")
	m.Server.OnSessionPacket = m.OnSessionPacket
	return m
}

func (mm *MgrMgr) Run() {
	go mm.Server.Run()
	go mm.Pm.Run()
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
