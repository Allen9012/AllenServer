package world

import (
	"github.com/Allen9012/AllenServer/demo/network"
	"github.com/Allen9012/AllenServer/demo/player"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	newPlayer := player.NewPlayer()
	newPlayer.UID = 111
	newPlayer.HandlerParamCh = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	newPlayer.Run()
}
