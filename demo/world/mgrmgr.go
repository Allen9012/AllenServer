package world

import (
	"github.com/Allen9012/AllenServer/demo/manager"
	"github.com/Allen9012/AllenServer/demo/network"
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
	Handlers        map[uint64]func(message *network.SessionPacket)
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
	if handler, ok := mm.Handlers[packet.Msg.ID]; ok {
		handler(packet)
	}
}
