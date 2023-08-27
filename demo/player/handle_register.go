package player

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (p *Player) HandlerRegister() {
	p.handlers[111] = p.AddFriend
	p.handlers[222] = p.DelFriend
	p.handlers[333] = p.HandleChatMsg
}
