package main

import "github.com/Allen9012/AllenServer/demo/network/protocol/gen/messageID"

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:
*/

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[messageID.MessageId_SCLogin] = c.OnLoginRsp
	c.messageHandlers[messageID.MessageId_SCAddFriend] = c.OnAddFriendRsp
	c.messageHandlers[messageID.MessageId_SCDelFriend] = c.OnDelFriendRsp
	c.messageHandlers[messageID.MessageId_SCSendChatMsg] = c.OnSendChatMsgRsp

}
