package module

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

// MgrInterface 管理器接口定义
type MgrInterface interface {
	OnStart()
	AfterStart()
	OnStop()
	AfterStop()
}

type Metrics interface {
	GetName() string
	SetName(str string)
}
