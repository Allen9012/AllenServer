package player

import "github.com/Allen9012/AllenServer/business/module/base"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

// Manager 维护在线玩家
type Manager struct {
	*base.MetricsBase
	players map[uint64]*Player
	addPCh  chan *Player
}

func (pm *Manager) OnStart() {
	//TODO implement me
	panic("implement me")
}

func (pm *Manager) AfterStart() {
	//TODO implement me
	panic("implement me")
}

func (pm *Manager) OnStop() {
	//TODO implement me
	panic("implement me")
}

func (pm *Manager) AfterStop() {
	//TODO implement me
	panic("implement me")
}
