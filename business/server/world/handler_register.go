package world

import "github.com/Allen9012/AllenServer/network/protocol/gen/messageID"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageID.MessageId_CSLogin] = mm.UserLogin
	mm.Handlers[messageID.MessageId_SCCreatePlayer] = mm.CreatePlayer
}
