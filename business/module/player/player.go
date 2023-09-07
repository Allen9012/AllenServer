package player

import (
	"github.com/Allen9012/AllenServer/business/module/chat"
	"github.com/Allen9012/AllenServer/business/module/friend"
	"github.com/Allen9012/AllenServer/business/module/task"
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
	HandlerParamCh chan *network.Message
	Session        *network.TcpConnX
	FriendSystem   *friend.System
	PrivateChat    *chat.PrivateChat
	taskData       *task.Data
}

// NewPlayer   TODO 分布式ID生成
func NewPlayer() *Player {
	p := &Player{
		UID:      0,
		taskData: task.NewTaskData(),
	}
	return p
}

// Start 不断处理从HandlerParamCh中的处理参数，然后取出消息来Handle
func (p *Player) Start() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			p.Handler(messageID.MessageId(handlerParam.ID), handlerParam)
		}
	}
}

func (p *Player) Stop() {

}

func (p *Player) OnLogin() {
	//从db加载数据初始化
	//同步数据给客户端
	p.taskData.LoadFromDB()
}

func (p *Player) OnLogout() {
	//存db
}

func (p *Player) GetTaskData() *task.Data {
	return p.taskData
}
