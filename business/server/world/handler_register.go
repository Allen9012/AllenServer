package world

import "github.com/Allen9012/AllenServer/network/protocol/gen/messageID"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (w *World) HandlerRegister() {
	w.Handlers[messageID.MessageId_CSCreatePlayer] = w.CreatePlayer
	w.Handlers[messageID.MessageId_CSLogin] = w.UserLogin
}
