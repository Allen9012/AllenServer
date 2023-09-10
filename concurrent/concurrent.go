package concurrent

import (
	"fmt"
	"github.com/Allen9012/AllenGame/log"
	"runtime"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
  @modified by:
**/

const defaultMaxTaskChannelNum = 1000000

/*
	cpuMul 表示cpu的倍数
	建议:(1)cpu密集型 使用1  (2)i/o密集型使用2或者更高
*/

type IConcurrent interface {
	OpenConcurrentByNumCPU(cpuMul float32)
	OpenConcurrent(minGoroutineNum int32, maxGoroutineNum int32, maxTaskChannelNum int)
	AsyncDoByQueue(queueId int64, fn func() bool, cb func(err error))
	AsyncDo(f func() bool, cb func(err error))
}

type Concurrent struct {
	dispatch // dispatch任务结构

	tasks     chan task
	cbChannel chan func(error) // 回调channel
}

/* =====Implement IConcurrent===== */

func (c *Concurrent) OpenConcurrentByNumCPU(cpuMul float32) {
	//TODO implement me
	panic("implement me")
}

func (c *Concurrent) OpenConcurrent(minGoroutineNum int32, maxGoroutineNum int32, maxTaskChannelNum int) {
	//TODO implement me
	panic("implement me")
}

func (c *Concurrent) AsyncDoByQueue(queueId int64, fn func() bool, cb func(err error)) {
	//TODO implement me
	panic("implement me")
}

func (c *Concurrent) AsyncDo(f func() bool, cb func(err error)) {
	//TODO implement me
	panic("implement me")
}

func (c *Concurrent) GetCallBackChannel() chan func(error) {
	return c.cbChannel
}

func (c *Concurrent) Close() {
	if cap(c.tasks) == 0 {
		return
	}

	log.Info("wait close concurrent")

	c.dispatch.close()

	log.Info("concurrent has successfully exited")
}

func (d *dispatch) DoCallback(cb func(err error)) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			l := runtime.Stack(buf, false)
			errString := fmt.Sprint(r)
			log.Dump(string(buf[:l]), log.String("error", errString))
		}
	}()

	cb(nil)
}
