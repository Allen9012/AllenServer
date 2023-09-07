package player

import (
	"github.com/Allen9012/AllenServer/business/module/chat"
	"github.com/Allen9012/AllenServer/business/module/friend"
	"github.com/Allen9012/AllenServer/business/module/task"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"google.golang.org/protobuf/proto"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

//var ERR_TYPE_CONVERT = errors.New("type convert error")

type Handler func(packet *network.Message)

func (p *Player) SendMsg(ID messageID.MessageId, message proto.Message) {
	id := uint64(ID)
	p.Session.AsyncSend(id, message)
}

func (p *Player) Handler(id messageID.MessageId, msg *network.Message) {
	if handler, _ := friend.GetHandler(id); handler != nil {
		handler.Fn(p.FriendSystem, msg)
	}
	if handler, _ := chat.GetHandler(id); handler != nil {
		handler.Fn(p.PrivateChat, msg)
	}
	if task.IsBelongToHere(id) {
		task.GetMe().ChIn <- &task.PlayerActionParam{
			MessageId: id,
			Player:    p,
			Packet:    msg,
		}
	}
}
