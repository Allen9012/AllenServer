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

type Manager struct {
	*base.MetricsBase
}

var (
	instance *Manager
	onceInit sync.Once
)

func GetMe() *Manager {
	onceInit.Do(func() {
		instance = &Manager{}
	})
	return instance
}

func (a *Manager) OnStart() {
	//TODO implement me
	panic("implement me")
}

func (a *Manager) AfterStart() {
	//TODO implement me
	panic("implement me")
}

func (a *Manager) OnStop() {
	//TODO implement me
	panic("implement me")
}

func (a *Manager) AfterStop() {
	//TODO implement me
	panic("implement me")
}
