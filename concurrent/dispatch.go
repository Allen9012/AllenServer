package concurrent

import (
	"github.com/Allen9012/AllenGame/util/queue"
	"sync"
	"sync/atomic"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc: dispatch
  @modified by:
**/

var idleTimeout = int64(2 * time.Second)

const maxTaskQueueSessionId = 10000

type dispatch struct {
	minConcurrentNum int32
	maxConcurrentNum int32

	queueIdChannel chan int64
	workerQueue    chan task
	tasks          chan task
	idle           bool
	workerNum      int32
	cbChannel      chan func(error) //异步任务队列

	mapTaskQueueSession map[int64]*queue.Deque[task]

	waitWorker   sync.WaitGroup
	waitDispatch sync.WaitGroup // 同步所有dispatch操作
}

func (d *dispatch) close() {
	//先置并发数为0
	atomic.StoreInt32(&d.minConcurrentNum, -1)
	//异步任务队列全部执行完成才可以最终close成功
breakFor:
	for {
		select {
		case cb := <-d.cbChannel:
			if cb == nil {
				break breakFor
			}
			cb(nil)
		}
	}

	d.waitDispatch.Wait()
}
