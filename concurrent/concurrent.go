package concurrent

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
  @modified by:
**/

const defaultMaxTaskChannelNum = 1000000

type IConcurrent interface {
	OpenConcurrentByNumCPU(cpuMul float32)
	OpenConcurrent(minGoroutineNum int32, maxGoroutineNum int32, maxTaskChannelNum int)
	AsyncDoByQueue(queueId int64, fn func() bool, cb func(err error))
	AsyncDo(f func() bool, cb func(err error))
}

type Concurrent struct {
	dispatch

	tasks     chan task
	cbChannel chan func(error)
}

func (c Concurrent) OpenConcurrentByNumCPU(cpuMul float32) {
	//TODO implement me
	panic("implement me")
}

func (c Concurrent) OpenConcurrent(minGoroutineNum int32, maxGoroutineNum int32, maxTaskChannelNum int) {
	//TODO implement me
	panic("implement me")
}

func (c Concurrent) AsyncDoByQueue(queueId int64, fn func() bool, cb func(err error)) {
	//TODO implement me
	panic("implement me")
}

func (c Concurrent) AsyncDo(f func() bool, cb func(err error)) {
	//TODO implement me
	panic("implement me")
}
