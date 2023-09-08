package service

import (
	"github.com/Allen9012/AllenGame/rpc"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: 服务模块
  @modified by:
**/

var timerDispatcherLen = 100000
var maxServiceEventChannelNum = 2000000

type IService interface {
	concurrent.IConcurrent
	Init(iService IService, getClientFun rpc.FuncRpcClient, getServerFun rpc.FuncRpcServer, serviceCfg interface{})
	Stop()
	Start()

	OnSetup(iService IService)
	OnInit() error
	OnStart()
	OnRelease()

	SetName(serviceName string)
	GetName() string
	GetRpcHandler() rpc.IRpcHandler
	GetServiceCfg() interface{}
	GetProfiler() *profiler.Profiler
	GetServiceEventChannelNum() int
	GetServiceTimerChannelNum() int

	SetEventChannelNum(num int)
	OpenProfiler()
}

type Service struct {
	Module

	rpcHandler             rpc.RpcHandler //rpc
	name                   string         //service name
	wg                     sync.WaitGroup
	serviceCfg             interface{}
	goroutineNum           int32
	startStatus            bool
	eventProcessor         event.IEventProcessor
	profiler               *profiler.Profiler //性能分析器
	nodeEventLister        rpc.INodeListener
	discoveryServiceLister rpc.IDiscoveryServiceListener
	chanEvent              chan event.IEvent
	closeSig               chan struct{}
}

// RpcConnEvent Node结点连接事件
type RpcConnEvent struct {
	IsConnect bool
	NodeId    int
}

// DiscoveryServiceEvent 发现服务结点
type DiscoveryServiceEvent struct {
	IsDiscovery bool
	ServiceName []string
	NodeId      int
}
