package main

import "github.com/Allen9012/AllenServer/demo/network/protocol/gen/messageID"

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:
*/

func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageID.MessageId_CSLogin.String()] = c.Login
	c.inputHandlers[messageID.MessageId_CSAddFriend.String()] = c.AddFriend
	c.inputHandlers[messageID.MessageId_CSDelFriend.String()] = c.DelFriend
	c.inputHandlers[messageID.MessageId_CSSendChatMsg.String()] = c.SendChatMsg
}

//login 10001 123456
//add_friend 10002
//del_friend 10003
