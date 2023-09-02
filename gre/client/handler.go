package main

import (
	"fmt"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/player"
	"google.golang.org/protobuf/proto"
	"strconv"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:
*/

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

func (c *Client) CreatePlayer(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &player.CSCreateUser{
		UserName: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)
}
func (c *Client) OnCreatePlayerRsp(packet *network.ClientPacket) {
	fmt.Println("恭喜你创建角色成功")
}

func (c *Client) Login(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &player.CSLogin{
		UserName: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)

}

func (c *Client) OnLoginRsp(packet *network.ClientPacket) {
	rsp := &player.SCLogin{}

	err := proto.Unmarshal(packet.Msg.Data, rsp)
	if err != nil {
		return
	}

	fmt.Println("登陆成功")
}

func (c *Client) AddFriend(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}

	uid, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSAddFriend{
		UID: uid,
	}
	c.Transport(id, msg)
}

func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {
	fmt.Println("add friend success !!")
}

func (c *Client) DelFriend(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}

	uid, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSDelFriend{
		UID: uid,
	}

	c.Transport(id, msg)
}

func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {
	fmt.Println("you have del friend success")

}

func (c *Client) SendChatMsg(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 3 { //""
		return
	}

	uid, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}
	category, err := strconv.ParseInt(param.Param[2], 10, 32)
	if err != nil {
		return
	}

	msg := &player.CSSendChatMsg{
		UID: uid,
		Msg: &player.ChatMessage{
			Content: param.Param[1],
			Extra:   nil,
		},
		Category: int32(category),
	}

	c.Transport(id, msg)
}

func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {
	fmt.Println("send  chat message success")
}
