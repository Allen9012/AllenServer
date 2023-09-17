package event

import "github.com/Allen9012/AllenGame/util/sync"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc: eventPool的内存池,缓存Event
  @modified by:
**/

const defaultMaxEventChannelNum = 2000000

var eventPool = sync.NewPoolEx(make(chan sync.IPoolData, defaultMaxEventChannelNum), func() sync.IPoolData {
	return &Event{}
})

func NewEvent() *Event {
	return eventPool.Get().(*Event)
}

func DeleteEvent(event IEvent) {
	eventPool.Put(event.(sync.IPoolData))
}

func SetEventPoolSize(eventPoolSize int) {
	eventPool = sync.NewPoolEx(make(chan sync.IPoolData, eventPoolSize), func() sync.IPoolData {
		return &Event{}
	})
}
