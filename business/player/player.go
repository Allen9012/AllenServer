package player

import (
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
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
	HandlerParamCh chan *network.Message
	handlers       map[messageID.MessageId]Handler
	Session        *network.Session
}

// NewPlayer   TODO 分布式ID生成
func NewPlayer() *Player {
	p := &Player{
		UID:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[messageID.MessageId]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[messageID.MessageId(handlerParam.ID)]; ok {
				fn(handlerParam)
			}
		}
	}
}

func (p *Player) OnLogin() {
	//从db加载数据初始化
	//同步数据给客户端

}

func (p *Player) OnLogout() {
	//存db
}
