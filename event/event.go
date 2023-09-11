package event

import "sync"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
  @modified by:
**/

// EventCallBack 事件接受器
type EventCallBack func(event IEvent)

type IEvent interface {
	GetEventType() EventType
}

type Event struct {
	Type EventType
	Data interface{}
	ref  bool
}

var emptyEvent Event

/*	=====Implement sync.IPoolData===== */
func (e *Event) Reset() {
	*e = emptyEvent
}

func (e *Event) IsRef() bool {
	return e.ref
}

func (e *Event) Ref() {
	e.ref = true
}

func (e *Event) UnRef() {
	e.ref = false
}

/*	=====Implement IEvent===== */
func (e *Event) GetEventType() EventType {
	return e.Type
}

// IEventHandler 事件Handler
type IEventHandler interface {
	Init(processor IEventProcessor)
	GetEventProcessor() IEventProcessor //获得事件
	NotifyEvent(IEvent)
	Destroy()
	//注册了事件
	addRegInfo(eventType EventType, eventProcessor IEventProcessor)
	removeRegInfo(eventType EventType, eventProcessor IEventProcessor)
}

type IEventChannel interface {
	PushEvent(ev IEvent) error
}

// IEventProcessor 事件处理程序
type IEventProcessor interface {
	IEventChannel

	Init(eventChannel IEventChannel)
	EventHandler(ev IEvent)
	RegEventReceiverFunc(eventType EventType, receiver IEventHandler, callback EventCallBack)
	UnRegEventReceiverFun(eventType EventType, receiver IEventHandler)

	castEvent(event IEvent) //广播事件
	addBindEvent(eventType EventType, receiver IEventHandler, callback EventCallBack)
	addListen(eventType EventType, receiver IEventHandler)
	removeBindEvent(eventType EventType, receiver IEventHandler)
	removeListen(eventType EventType, receiver IEventHandler)
}

/*	Implement IEventHandler  */
type EventHandler struct {
	//已经注册的事件类型
	eventProcessor IEventProcessor

	//已经注册的事件
	locker      sync.RWMutex
	mapRegEvent map[EventType]map[IEventProcessor]interface{} //向其他事件处理器监听的事件类型
}

/*	Implement IEventProcessor  */
type EventProcessor struct {
	IEventChannel

	locker              sync.RWMutex
	mapListenerEvent    map[EventType]map[IEventProcessor]int         //监听者信息
	mapBindHandlerEvent map[EventType]map[IEventHandler]EventCallBack //收到事件处理和回调的map
}

/*	=====Implement IEventProcessor=====  */

func (e *EventProcessor) Init(eventChannel IEventChannel) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) EventHandler(ev IEvent) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) RegEventReceiverFunc(eventType EventType, receiver IEventHandler, callback EventCallBack) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) UnRegEventReceiverFun(eventType EventType, receiver IEventHandler) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) castEvent(event IEvent) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) addBindEvent(eventType EventType, receiver IEventHandler, callback EventCallBack) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) addListen(eventType EventType, receiver IEventHandler) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) removeBindEvent(eventType EventType, receiver IEventHandler) {
	//TODO implement me
	panic("implement me")
}

func (e *EventProcessor) removeListen(eventType EventType, receiver IEventHandler) {
	//TODO implement me
	panic("implement me")
}

/*	=====Implement IEventHandler===== */

func (e *EventHandler) Init(processor IEventProcessor) {
	//TODO implement me
	panic("implement me")
}

func (e *EventHandler) GetEventProcessor() IEventProcessor {
	//TODO implement me
	panic("implement me")
}

func (e *EventHandler) NotifyEvent(event IEvent) {
	//TODO implement me
	panic("implement me")
}

func (e *EventHandler) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (e *EventHandler) addRegInfo(eventType EventType, eventProcessor IEventProcessor) {
	//TODO implement me
	panic("implement me")
}

func (e *EventHandler) removeRegInfo(eventType EventType, eventProcessor IEventProcessor) {
	//TODO implement me
	panic("implement me")
}

func NewEventProcessor() IEventProcessor {
	ep := EventProcessor{}
	ep.mapListenerEvent = map[EventType]map[IEventProcessor]int{}
	ep.mapBindHandlerEvent = map[EventType]map[IEventHandler]EventCallBack{}

	return &ep
}
