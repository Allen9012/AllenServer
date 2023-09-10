package rpc

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
  @modified by:
**/

type CallTimer struct {
	SeqId    uint64
	FireTime int64
}

type CallTimerHeap struct {
	callTimer   []CallTimer
	mapSeqIndex map[uint64]int
}
