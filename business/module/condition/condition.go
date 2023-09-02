package condition

import "github.com/Allen9012/AllenServer/business/module/condition/event"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type Condition interface {
	CheckArrived() bool
	OnNotify(event.Event)
	GetId() uint32
	SetCB(func())
}

type Base struct {
	Cb func()
}

func NewTargetBase() *Base {
	return &Base{}
}

func (t *Base) CheckArrived() bool {
	return false
}

func (t *Base) OnNotify(event event.Event) {

}

func (t Base) GetId() uint32 {
	return 0
}

func (t *Base) SetCB(f func()) {
	t.Cb = f
}
