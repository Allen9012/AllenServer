package player

import (
	"github.com/Allen9012/AllenServer/demo/define"
	"github.com/Allen9012/AllenServer/demo/network"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

// Player 玩家
type Player struct {
	UID            uint64
	FriendList     []uint64 //朋友
	HandlerParamCh chan *define.HandlerParam
	handlers       map[uint64]Handler
	session        *network.Session
}

// NewPlayer   TODO 分布式ID生成
func NewPlayer() *Player {
	p := &Player{
		UID:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[uint64]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandleParamCh:
			if fn, ok := p.handlers[handlerParam.ID]; ok {
				fn(handlerParam.Data)
			}
		}
	}
}
