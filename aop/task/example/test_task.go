package example

import "github.com/Allen9012/AllenServer/aop/task"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type TTask struct {
	Conf    *task.Config
	Next    *TTask
	Status  task.Status
	Targets []task.Target
}

func NewTTask(config *task.Config) *TTask {
	t := &TTask{
		Conf: config,
	}
	return t

}

func (t *TTask) Accept(config *task.Config) {
	t.Status = task.ACCEPT
}

func (t *TTask) Finish() {
	t.Status = task.FINISH

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
