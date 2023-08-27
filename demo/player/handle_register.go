package player

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (p *Player) HandleRegister() {
	p.handlers["add_friend"] = p.AddFriend
	p.handlers["del_friend"] = p.DelFriend
	p.handlers["handle_chat_msg"] = p.HandleChatMsg
}
