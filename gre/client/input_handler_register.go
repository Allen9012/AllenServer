package main

import (
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"google.golang.org/protobuf/proto"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:
*/

func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageID.MessageId_CSCreatePlayer.String()] = c.CreatePlayer
	c.inputHandlers[messageID.MessageId_CSLogin.String()] = c.Login
	c.inputHandlers[messageID.MessageId_CSAddFriend.String()] = c.AddFriend
	c.inputHandlers[messageID.MessageId_CSDelFriend.String()] = c.DelFriend
	c.inputHandlers[messageID.MessageId_CSSendChatMsg.String()] = c.SendChatMsg
}

// GetMessageIdByCmd cmd字符串转化为常量
func (c *Client) GetMessageIdByCmd(cmd string) messageID.MessageId {
	mid, ok := messageID.MessageId_value[cmd]
	if ok {
		return messageID.MessageId(mid)
	}
	return messageID.MessageId_None
}

// Transport 最后加密消息和发送消息
func (c *Client) Transport(id messageID.MessageId, message proto.Message) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	c.cli.ChMsg <- &network.Message{
		ID:   uint64(id),
		Data: bytes,
	}
}
