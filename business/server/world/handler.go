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

func (w *World) CreatePlayer(message *network.Packet) {
	msg := &player.CSCreateUser{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	fmt.Println("[World.CreatePlayer]", msg)
	w.SendMsg(uint64(messageID.MessageId_SCCreatePlayer), &player.SCCreateUser{}, message.Conn)

}

func (w *World) UserLogin(message *network.Packet) {
	msg := &player.CSLogin{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	newPlayer := logicPlayer.NewPlayer()
	// TODO ID生成
	newPlayer.UID = 111
	newPlayer.Session = message.Conn
	w.Pm.Add(newPlayer)

}

func (w *World) SendMsg(id uint64, message proto.Message, session *network.TcpConnX) {
	session.AsyncSend(id, message)
}
