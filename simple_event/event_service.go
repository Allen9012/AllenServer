package simple_event

import (
	"fmt"
	"github.com/Allen9012/AllenGame/event"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
	"github.com/Allen9012/AllenGame/util/timer"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/11
  @desc: event service事件服务
  @modified by:
**/

func init() {
	node.Setup(&ServiceEvent{}, &ServiceEventReceive{})
}

// 发送消息服务
type ServiceEvent struct {
	service.Service
}

func (slf *ServiceEvent) OnInit() error {
	fmt.Printf("【事件服务】启动\n")

	slf.AfterFunc(time.Second*1, slf.TriggerEvent)

	return nil
}

// 广播消息
func (slf *ServiceEvent) TriggerEvent(t *timer.Timer) {
	slf.GetEventHandler().NotifyEvent(&event.Event{
		Type: SimpleEvent1,
		Data: "自定义事件1",
	})
	slf.AfterFunc(time.Second*1, slf.TriggerEvent)
}

// -------------------------------------------------------------------------------------

// 接受消息服务
type ServiceEventReceive struct {
	service.Service
}

func (slf *ServiceEventReceive) OnInit() error {
	fmt.Printf("【接受消息服务】启动\n")

	pService := node.GetService("ServiceEvent")

	pService.(*ServiceEvent).GetEventHandler().GetEventProcessor().RegEventReceiverFunc(SimpleEvent1, slf.GetEventHandler(), slf.OnEvent)

	slf.AddModule(&EventModule{})

	return nil
}

func (slf *ServiceEventReceive) OnEvent(ev event.IEvent) {
	event := ev.(*event.Event)
	fmt.Printf("ServiceEventReceive 收到事件 event: %v\n", event)
}

func (slf *ServiceEventReceive) OnRelease() {
	pService := node.GetService("ServiceEvent")

	pService.(*ServiceEvent).GetEventHandler().GetEventProcessor().UnRegEventReceiverFun(SimpleEvent1, slf.GetEventHandler())
}

//---------------------------------------------------------------------------------

type EventModule struct {
	service.Module
}

func (slf *EventModule) OnInit() error {

	pService := node.GetService("ServiceEvent")

	pService.(*ServiceEvent).GetEventHandler().GetEventProcessor().RegEventReceiverFunc(SimpleEvent1, slf.GetEventHandler(), slf.OnEvent)

	return nil
}

func (slf *EventModule) OnEvent(ev event.IEvent) {
	event := ev.(*event.Event)
	fmt.Printf("EventModule 收到事件 event: %v\n", event)
}

func (slf *EventModule) OnRelease() {
	pService := node.GetService("ServiceEvent")

	pService.(*ServiceEvent).GetEventHandler().GetEventProcessor().UnRegEventReceiverFun(SimpleEvent1, slf.GetEventHandler())
}
