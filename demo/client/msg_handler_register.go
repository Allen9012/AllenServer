package main

/*
*

	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:

*
*/
func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[111] = c.OnLoginRsp
	c.messageHandlers[222] = c.OnAddFriendRsp
	c.messageHandlers[333] = c.OnDelFriendRsp
	c.messageHandlers[444] = c.OnSendChatMsgRsp

}
