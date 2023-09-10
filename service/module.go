package service

import (
	"github.com/Allen9012/AllenGame/concurrent"
	"github.com/Allen9012/AllenGame/event"
	rpcHandle "github.com/Allen9012/AllenGame/rpc"
	"github.com/Allen9012/AllenGame/util/timer"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
  @modified by:
**/

const InitModuleId = 1e9

type IModule interface {
	concurrent.IConcurrent
	SetModuleId(moduleId uint32) bool
	GetModuleId() uint32
	AddModule(module IModule) (uint32, error)
	GetModule(moduleId uint32) IModule
	GetAncestor() IModule
	ReleaseModule(moduleId uint32)
	NewModuleId() uint32
	GetParent() IModule
	OnInit() error
	OnRelease()
	getBaseModule() IModule
	GetService() IService
	GetModuleName() string
	GetEventProcessor() event.IEventProcessor
	NotifyEvent(ev event.IEvent)
}

type IModuleTimer interface {
	AfterFunc(d time.Duration, cb func(*timer.Timer)) *timer.Timer
	CronFunc(cronExpr *timer.CronExpr, cb func(*timer.Cron)) *timer.Cron
	NewTicker(d time.Duration, cb func(*timer.Ticker)) *timer.Ticker
}

type Module struct {
	rpcHandle.IRpcHandler
	moduleId         uint32             //模块Id
	moduleName       string             //模块名称
	parent           IModule            //父亲
	self             IModule            //自己
	child            map[uint32]IModule //孩子们
	mapActiveTimer   map[timer.ITimer]struct{}
	mapActiveIdTimer map[uint64]timer.ITimer
	dispatcher       *timer.Dispatcher //timer

	//根结点
	ancestor     IModule            //始祖
	seedModuleId uint32             //模块id种子
	descendants  map[uint32]IModule //始祖的后裔们

	//事件管道
	eventHandler event.IEventHandler
	concurrent.IConcurrent
}
