package player

import (
	"fmt"
	"github.com/Allen9012/AllenServer/demo/function"
	"github.com/Allen9012/AllenServer/demo/network"
	"github.com/Allen9012/AllenServer/demo/network/protocol/gen/player"
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
	if err := proto.Unmarshal(packet.Data, req); err != nil {
		fmt.Println("CSAddFriend Unmarshal error")
		return
	}
	if !function.CheckInNumberSlice(req.UID, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UID)
	}
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
	if err := proto.Unmarshal(packet.Data, req); err != nil {
		fmt.Println("CSAddFriend Unmarshal error")
		return
	}
	p.FriendList = function.DelEleInSlice(req.UID, p.FriendList)
}

func (p *Player) HandleChatMsg(packet *network.Message) {
	req := &player.CSSendChatMsg{}
	if err := proto.Unmarshal(packet.Data, req); err != nil {
		fmt.Println("CSAddFriend Unmarshal error")
		return
	}
	fmt.Println(req.Msg.Content)
	//	todo， 收到消息，传送给客户端
}
