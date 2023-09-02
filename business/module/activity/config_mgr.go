package activity

import (
	"github.com/Allen9012/AllenServer/business/module/base"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type ConfigManager struct {
	base.ConfigManagerBase
	Configs sync.Map
}

func (m *ConfigManager) Init(id uint32) interface{} {
	var ret any
	m.Configs.Range(func(key, value any) bool {
		idAssert := key.(uint32)
		if idAssert == id {
			ret = value
			return false
		}
		return true
	})
	return ret
}
