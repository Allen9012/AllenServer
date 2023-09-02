package example

import (
	"fmt"
	"github.com/Allen9012/AllenServer/aop/task"
	"testing"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/9/1
	@desc:
	@modified by:
*/

func TestName(t *testing.T) {
	te := TEvent{
		Subscribers: make([]task.Target, 0),
	}
	tg := &TTarget{
		Id:   111,
		Data: 1,
	}
	te.Attach(tg)
	te.Data = 1
	te.Notify()
	fmt.Println("CheckDone:", tg.CheckDone())
}
