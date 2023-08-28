package main

import (
	"fmt"
	"github.com/Allen9012/AllenServer/demo/network"
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

func (c *Client) Login(param *InputParam) {
	fmt.Println("login input handler")
	fmt.Println(param.Command)
	fmt.Println(param.Param)
}

func (c *Client) OnLoginRsp(packet *network.ClientPacket) {

}

func (c *Client) AddFriend(param *InputParam) {

}

func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {

}

func (c *Client) DelFriend(param *InputParam) {

}

func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {

}

func (c *Client) SendChatMsg(param *InputParam) {

}

func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {

}
