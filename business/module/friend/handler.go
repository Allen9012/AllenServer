package friend

import (
	"errors"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"github.com/Allen9012/AllenServer/network/protocol/gen/player"
	"github.com/Allen9012/sugar"
	"google.golang.org/protobuf/proto"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/3
  @desc:
  @modified by:
**/

type Handler struct {
	Id messageID.MessageId
	Fn func(s *System, packet *network.Message)
}

var (
	handlers     []*Handler
	onceInit     sync.Once
	MinMessageId messageID.MessageId
	MaxMessageId messageID.MessageId //handle 的消息范围
)

func IsBelongToHere(id messageID.MessageId) bool {
	return id > MinMessageId && id < MaxMessageId
}

func GetHandler(id messageID.MessageId) (*Handler, error) {

	if id > MinMessageId && id < MaxMessageId {
		return nil, errors.New("not in")
	}
	for _, handler := range handlers {
		if handler.Id == id {
			return handler, nil
		}
	}
	return nil, errors.New("not exist")
}

func init() {
	onceInit.Do(func() {
		HandlerFriendRegister()
	})
}

func HandlerFriendRegister() {
	handlers[0] = &Handler{
		messageID.MessageId_CSAddFriend,
		AddFriend,
	}
	handlers[1] = &Handler{
		messageID.MessageId_CSDelFriend,
		DelFriend,
	}
}

func GetFriendList(s *System, packet *network.Message) {

}

func GetFriendInfo(s *System, packet *network.Message) {

}

func AddFriend(s *System, packet *network.Message) {
	req := &player.CSAddFriend{}

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}

	if !sugar.CheckInSlice(req.UID, s.FriendList) {
		s.FriendList = append(s.FriendList, req.UID)
	}
	s.Owner.SendMsg(messageID.MessageId_SCAddFriend, &player.SCSendChatMsg{})

}

func DelFriend(s *System, packet *network.Message) {
	req := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	s.FriendList = sugar.DelOneInSlice(req.UID, s.FriendList)

	s.Owner.SendMsg(messageID.MessageId_SCDelFriend, &player.SCDelFriend{})
}

func GiveFriendItem(s *System, packet *network.Message) {

}

func AddApply(s *System, packet *network.Message) {

}

func ManagerApply(s *System, packet *network.Message) {

}
