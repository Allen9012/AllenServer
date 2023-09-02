package example

import (
	task2 "github.com/Allen9012/AllenServer/business/module/task"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type TTask struct {
	Conf    *task2.Config
	Next    *TTask
	Status  task2.Status
	Targets []task2.Target
}

func NewTTask(config *task2.Config) *TTask {
	t := &TTask{
		Conf: config,
	}
	return t

}

func (t *TTask) Accept(config *task2.Config) {
	t.Status = task2.ACCEPT
}

func (t *TTask) Finish() {
	t.Status = task2.FINISH

}

func (t *TTask) TargetDoneCallBack() {
	count := 0
	for _, target := range t.Targets {
		if target.CheckDone() {
			count++
		}
	}
	if count == len(t.Targets) {
		t.Finish()
	}
}