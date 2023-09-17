package event

import (
	"fmt"
	"github.com/Allen9012/AllenGame/log"
	"runtime"
	"sync"
)

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

func NewEventHandler() IEventHandler {
	eh := EventHandler{}
	eh.mapRegEvent = map[EventType]map[IEventProcessor]interface{}{}

	return &eh
}

func NewEventProcessor() IEventProcessor {
	ep := EventProcessor{}
	ep.mapListenerEvent = map[EventType]map[IEventProcessor]int{}
	ep.mapBindHandlerEvent = map[EventType]map[IEventHandler]EventCallBack{}

	return &ep
}

/*	=====EventProcessor Implement IEventProcessor=====  */

func (processor *EventProcessor) Init(eventChannel IEventChannel) {
	processor.IEventChannel = eventChannel
}

func (processor *EventProcessor) EventHandler(ev IEvent) {
	// 事件的处理需要recover
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			l := runtime.Stack(buf, false)
			errString := fmt.Sprint(r)
			log.Dump(string(buf[:l]), log.String("error", errString))
		}
	}()
	mapCallBack, ok := processor.mapBindHandlerEvent[ev.GetEventType()]
	if !ok {
		return
	}
	// 执行callback
	for _, callback := range mapCallBack {
		callback(ev)
	}
}

// 把事件注册到对应的Handler
func (processor *EventProcessor) RegEventReceiverFunc(eventType EventType, receiver IEventHandler, callback EventCallBack) {
	//记录receiver自己注册过的事件
	receiver.addRegInfo(eventType, processor)
	//记录当前所属IEventProcessor注册的回调
	receiver.GetEventProcessor().addBindEvent(eventType, receiver, callback)
	//将注册加入到监听中
	processor.addListen(eventType, receiver)
}

func (processor *EventProcessor) UnRegEventReceiverFun(eventType EventType, receiver IEventHandler) {
	processor.removeListen(eventType, receiver)
	receiver.GetEventProcessor().removeBindEvent(eventType, receiver)
	receiver.removeRegInfo(eventType, processor)
}

// 事件分发到对应的Service
func (processor *EventProcessor) castEvent(event IEvent) {
	if processor.mapListenerEvent == nil {
		log.Error("mapListenerEvent not init!")
		return
	}
	eventProcessor, ok := processor.mapListenerEvent[event.GetEventType()]
	if !ok {
		log.Debug("event is not listen", log.Int("event type", int(event.GetEventType())))
		return
	}
	for proc := range eventProcessor {
		proc.PushEvent(event)
	}
}

// 增加对应Handler的callback
func (processor *EventProcessor) addBindEvent(eventType EventType, receiver IEventHandler, callback EventCallBack) {
	processor.locker.Lock()
	defer processor.locker.Unlock()
	if _, ok := processor.mapBindHandlerEvent[eventType]; !ok {
		// 如果已有就更新
		processor.mapBindHandlerEvent[eventType] = map[IEventHandler]EventCallBack{}
	}

	processor.mapBindHandlerEvent[eventType][receiver] = callback
}

func (processor *EventProcessor) addListen(eventType EventType, receiver IEventHandler) {
	processor.locker.Lock()
	defer processor.locker.Unlock()
	if _, ok := processor.mapListenerEvent[eventType]; !ok {
		// 如没有就初始化
		processor.mapListenerEvent[eventType] = map[IEventProcessor]int{}
	}

	processor.mapListenerEvent[eventType][receiver.GetEventProcessor()] += 1
}

func (processor *EventProcessor) removeBindEvent(eventType EventType, receiver IEventHandler) {
	processor.locker.Lock()
	defer processor.locker.Unlock()
	if _, ok := processor.mapBindHandlerEvent[eventType]; ok == true {
		delete(processor.mapBindHandlerEvent[eventType], receiver)
	}
}

func (processor *EventProcessor) removeListen(eventType EventType, receiver IEventHandler) {
	processor.locker.Lock()
	defer processor.locker.Unlock()
	if _, ok := processor.mapListenerEvent[eventType]; ok == true {
		processor.mapListenerEvent[eventType][receiver.GetEventProcessor()] -= 1
		if processor.mapListenerEvent[eventType][receiver.GetEventProcessor()] <= 0 {
			delete(processor.mapListenerEvent[eventType], receiver.GetEventProcessor())
		}
	}
}

/*	=====EventHandler Implement IEventHandler===== */

func (handler *EventHandler) Init(processor IEventProcessor) {
	handler.eventProcessor = processor
	handler.mapRegEvent = map[EventType]map[IEventProcessor]interface{}{}
}

func (handler *EventHandler) GetEventProcessor() IEventProcessor {
	return handler.eventProcessor
}

func (handler *EventHandler) NotifyEvent(event IEvent) {
	handler.GetEventProcessor().castEvent(event)
}

// 清除所有mapEventProcess
func (handler *EventHandler) Destroy() {
	handler.locker.Lock()
	defer handler.locker.Unlock()
	for eventTyp, mapEventProcess := range handler.mapRegEvent {
		if mapEventProcess == nil {
			continue
		}

		for eventProcess := range mapEventProcess {
			eventProcess.UnRegEventReceiverFun(eventTyp, handler)
		}
	}
}

func (handler *EventHandler) addRegInfo(eventType EventType, eventProcessor IEventProcessor) {
	handler.locker.Lock()
	defer handler.locker.Unlock()
	// 如果还没有，先初始化
	if handler.mapRegEvent == nil {
		handler.mapRegEvent = map[EventType]map[IEventProcessor]interface{}{}
	}
	if _, ok := handler.mapRegEvent[eventType]; !ok {
		handler.mapRegEvent[eventType] = map[IEventProcessor]interface{}{}
	}
	handler.mapRegEvent[eventType][eventProcessor] = nil
}

// 删除eventProcessor 的信息
func (handler *EventHandler) removeRegInfo(eventType EventType, eventProcessor IEventProcessor) {
	if _, ok := handler.mapRegEvent[eventType]; ok == true {
		delete(handler.mapRegEvent[eventType], eventProcessor)
	}
}
