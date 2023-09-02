package manager

import (
	"github.com/Allen9012/AllenServer/business/player"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

// PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]*player.Player
	addPCh  chan *player.Player
}

// Add 玩家加入玩家组
func (pm *PlayerMgr) Add(p *player.Player) {
	if pm.players[p.UID] != nil {
		return
	}
	pm.players[p.UID] = p
	go p.Run()
}

// Del ...
func (pm *PlayerMgr) Del(p player.Player) {
	delete(pm.players, p.UID)
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		}
	}
}

func (pm *PlayerMgr) GetPlayer(uId uint64) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
