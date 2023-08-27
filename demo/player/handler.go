package player

import (
	"errors"
	"fmt"
	"github.com/Allen9012/AllenServer/demo/chat"
	"github.com/Allen9012/AllenServer/utils/function"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

var ERR_TYPE_CONVERT = errors.New("type convert error")

type Handler func(data interface{})

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

func (p *Player) AddFriend(data interface{}) {
	FriendID := data.(uint64)
	if !function.CheckInNumberSlice(FriendID, p.FriendList) {
		p.FriendList = append(p.FriendList, FriendID)
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

func (p *Player) DelFriend(data interface{}) {
	FriendID := data.(uint64)
	p.FriendList = function.DelEleInSlice(FriendID, p.FriendList)
}

func (p *Player) HandleChatMsg(data interface{}) {
	chatMsg := data.(chat.Msg)
	fmt.Println(chatMsg)
	//	todo， 收到消息，传送给客户端
}
