package concurrent

import (
	"github.com/Allen9012/AllenGame/util/queue"
	"sync"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
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
	cbChannel      chan func(error)

	mapTaskQueueSession map[int64]*queue.Deque[task]

	waitWorker   sync.WaitGroup
	waitDispatch sync.WaitGroup
}
