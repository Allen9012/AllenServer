package task

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type Task interface {
	Accept(config *Config)
	Finish()
	TargetDoneCallBack()
}

type Base struct {
}

func (b *Base) Accept(config *Config) {

}

func (b *Base) Finish() {

}

func (b *Base) TargetDoneCallBack() {

}
