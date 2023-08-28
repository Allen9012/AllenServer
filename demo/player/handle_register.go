package player

import "github.com/Allen9012/AllenServer/demo/network/protocol/gen/messageID"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (p *Player) HandlerRegister() {
	p.handlers[messageID.MessageId_SCAddFriend] = p.AddFriend
	p.handlers[messageID.MessageId_SCDelFriend] = p.DelFriend
	p.handlers[messageID.MessageId_SCSendChatMsg] = p.HandleChatMsg
}
