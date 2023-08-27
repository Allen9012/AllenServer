package player

import "github.com/Allen9012/AllenServer/demo/define"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

// Player 玩家
type Player struct {
	UID           uint64
	FriendList    []uint64 //朋友
	HandleParamCh chan *define.HandlerParam
	handlers      map[string]Handler
}

func NewPlayer(uid uint64) *Player {
	p := &Player{
		UID:        0,
		FriendList: nil,
	}
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandleParamCh:
			if fn, ok := p.handlers[handlerParam.HandlerKey]; ok {
				fn(handlerParam.Data)
			}
		}
	}
}
