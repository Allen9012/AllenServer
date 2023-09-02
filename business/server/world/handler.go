package world

import (
	"fmt"
	logicPlayer "github.com/Allen9012/AllenServer/business/module/player"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"github.com/Allen9012/AllenServer/network/protocol/gen/player"
	"google.golang.org/protobuf/proto"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	msg := &player.CSLogin{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	newPlayer := logicPlayer.NewPlayer()
	newPlayer.UID = 111
	//newPlayer.UID = uint64(time.Now().Unix())
	newPlayer.HandlerParamCh = message.Sess.WriteCh
	message.Sess.IsPlayerOnline = true
	message.Sess.UID = newPlayer.UID
	newPlayer.Session = message.Sess
	mm.Pm.Add(newPlayer)
}

func (mm *MgrMgr) CreatePlayer(message *network.SessionPacket) {
	msg := &player.CSCreateUser{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	//TODO 存储逻辑
	fmt.Println("[MgrMgr.CreatePlayer]", msg)
	// 回复创角消息
	mm.SendMsg(uint64(messageID.MessageId_SCCreatePlayer), &player.SCCreateUser{}, message.Sess)
}

func (mm *MgrMgr) SendMsg(id uint64, message proto.Message, session *network.Session) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	rsp := &network.Message{
		ID:   id,
		Data: bytes,
	}
	session.SendMsg(rsp)
}
