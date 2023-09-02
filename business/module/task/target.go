package task

import (
	"github.com/Allen9012/AllenServer/business/module/condition/event"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type Target interface {
	CheckDone() bool
	OnNotify(event.Event)
	GetTargetId() uint32
	SetTaskCB(func())
}

type TargetBase struct {
	TaskCB func()
}

func NewTargetBase() *TargetBase {
	return &TargetBase{}
}

func (t *TargetBase) CheckDone() bool {
	return false
}

func (t *TargetBase) OnNotify(event event.Event) {

}

func (t *TargetBase) GetTargetId() uint32 {
	return 0
}

func (t *TargetBase) SetTaskCB(f func()) {
	t.TaskCB = f
}
