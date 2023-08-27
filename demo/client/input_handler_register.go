package main

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:
*/

func (c *Client) InputHandlerRegister() {
	c.inputHandlers["login"] = c.Login
	c.inputHandlers["add_friend"] = c.AddFriend
	c.inputHandlers["del_friend"] = c.DelFriend
	c.inputHandlers["chat_msg"] = c.SendChatMsg
}
