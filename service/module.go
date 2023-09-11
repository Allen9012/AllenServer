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

func (m *Module) SetModuleId(moduleId uint32) bool {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetModuleId() uint32 {
	//TODO implement me
	panic("implement me")
}

func (m *Module) AddModule(module IModule) (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetModule(moduleId uint32) IModule {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetAncestor() IModule {
	//TODO implement me
	panic("implement me")
}

func (m *Module) ReleaseModule(moduleId uint32) {
	//TODO implement me
	panic("implement me")
}

func (m *Module) NewModuleId() uint32 {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetParent() IModule {
	//TODO implement me
	panic("implement me")
}

func (m *Module) OnInit() error {
	//TODO implement me
	panic("implement me")
}

// OnRelease 释放和关闭service需要释放Module
func (m *Module) OnRelease() {
}

func (m *Module) getBaseModule() IModule {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetService() IService {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetModuleName() string {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetEventProcessor() event.IEventProcessor {
	//TODO implement me
	panic("implement me")
}

func (m *Module) NotifyEvent(ev event.IEvent) {
	//TODO implement me
	panic("implement me")
}
