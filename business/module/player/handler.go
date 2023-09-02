package player

import (
	"fmt"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"github.com/Allen9012/AllenServer/network/protocol/gen/player"
	"github.com/Allen9012/sugar"
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

func (p *Player) AddFriend(packet *network.Message) {
	req := &player.CSAddFriend{}

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}

	if !sugar.CheckInSlice(req.UID, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UID)
	}

	p.SendMsg(messageID.MessageId_SCAddFriend, &player.SCSendChatMsg{})
}

func (p *Player) DelFriend(packet *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	p.FriendList = sugar.DelOneInSlice(req.UID, p.FriendList)

	p.SendMsg(messageID.MessageId_SCDelFriend, &player.SCDelFriend{})
}

func (p *Player) HandleChatMsg(packet *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)

	p.SendMsg(messageID.MessageId_SCSendChatMsg, &player.SCSendChatMsg{})
}

func (p *Player) SendMsg(ID messageID.MessageId, message proto.Message) {
	id := uint64(ID)
	p.Session.AsyncSend(uint16(id), message)
}
