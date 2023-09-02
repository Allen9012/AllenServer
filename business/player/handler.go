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

//func (p *Player) AddFriend(data interface{}) (bool, error) {
//	FriendID, ok := data.(uint64)
//	if !ok {
//		fmt.Println("AddFriend data error")
//		return false, ERR_TYPE_CONVERT
//	}
//	if !function.CheckInNumberSlice(FriendID, p.FriendList) {
//		p.FriendList = append(p.FriendList, FriendID)
//		return true, nil
//	}
//	return false, nil
//}

func (p *Player) AddFriend(packet *network.Message) {
	req := &player.CSAddFriend{}

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}

	if !sugar.CheckInSlice(req.UID, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UID)
	}

	bytes, err := proto.Marshal(&player.SCSendChatMsg{})
	if err != nil {
		return
	}

	rsp := &network.Message{
		ID:   uint64(messageID.MessageId_SCAddFriend),
		Data: bytes,
	}

	p.Session.SendMsg(rsp)
}

//func (p *Player) DelFriend(data interface{}) (bool, error) {
//	FriendID, ok := data.(uint64)
//	if !ok {
//		fmt.Println("AddFriend data error")
//		return false, ERR_TYPE_CONVERT
//	}
//	p.FriendList = function.DelEleInSlice(FriendID, p.FriendList)
//	return true, nil
//}

func (p *Player) DelFriend(packet *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	p.FriendList = sugar.DelOneInSlice(req.UID, p.FriendList)

	bytes, err := proto.Marshal(&player.SCDelFriend{})
	if err != nil {
		return
	}

	rsp := &network.Message{
		ID:   uint64(messageID.MessageId_SCDelFriend),
		Data: bytes,
	}

	p.Session.SendMsg(rsp)
}

func (p *Player) HandleChatMsg(packet *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)

	bytes, err := proto.Marshal(&player.SCSendChatMsg{})
	if err != nil {
		return
	}

	rsp := &network.Message{
		ID:   uint64(messageID.MessageId_SCSendChatMsg),
		Data: bytes,
	}

	p.Session.SendMsg(rsp)
}
