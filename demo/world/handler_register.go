package world

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[1] = mm.UserLogin
}
