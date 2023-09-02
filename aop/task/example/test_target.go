package example

import "github.com/Allen9012/AllenServer/aop/task"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type TTarget struct {
	Id   uint32
	Data int
	Done bool
}

func NewTTarget() {

}

func (T TTarget) CheckDone() bool {
	return T.Done
}

func (T *TTarget) OnNotify(event task.Event) {
	e := event.(*TEvent)
	if e.Data == T.Data {
		T.Done = true
	}
}

func (T TTarget) GetTargetId() uint32 {
	return T.Id
}

func (T TTarget) SetTaskCB(fn func()) {

}
