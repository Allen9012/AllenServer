package mgrmgr

import (
	"github.com/Allen9012/AllenServer/demo/manager"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 保存所有的manager
  @modified by:
**/

type MgrMgr struct {
	Pm manager.PlayerMgr
}

var MM *MgrMgr

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
	return m
}
