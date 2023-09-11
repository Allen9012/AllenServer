package service

import (
	"fmt"
	"github.com/Allen9012/AllenGame/concurrent"
	"github.com/Allen9012/AllenGame/event"
	"github.com/Allen9012/AllenGame/log"
	"github.com/Allen9012/AllenGame/profiler"
	"github.com/Allen9012/AllenGame/rpc"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
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
	startStatus            bool //状态位
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

func (s *Service) Init(iService IService, getClientFun interface{}, getServerFun interface{}, serviceCfg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Stop() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Start() {
	s.startStatus = true
	var waitRun sync.WaitGroup

	for i := int32(0); i < s.goroutineNum; i++ {
		s.wg.Add(1)
		waitRun.Add(1)
		go func() {
			log.Info(s.GetName() + " service is running")
			waitRun.Done()
			s.Run()
		}()
	}

	waitRun.Wait()
}

// Run Service运行
func (s *Service) Run() {
	defer s.wg.Done()
	var bStop = false
	// 拿到并发回调channel
	concur := s.IConcurrent.(*concurrent.Concurrent)
	concurrentCBChannel := concur.GetCallBackChannel()

	s.self.(IService).OnStart()
	for {
		var analyzer *profiler.Analyzer
		select {
		case <-s.closeSig:
			bStop = true
			concur.Close()
		case cb := <-concurrentCBChannel:
			concur.DoCallback(cb)
		case ev := <-s.chanEvent:
			switch ev.GetEventType() {
			case event.ServiceRpcRequestEvent:
				cEvent, ok := ev.(*event.Event)
				if ok == false {
					log.Error("Type event conversion error")
					break
				}
				rpcRequest, ok := cEvent.Data.(*rpc.RpcRequest)
				if ok == false {
					log.Error("Type *rpc.RpcRequest conversion error")
					break
				}
				if s.profiler != nil {
					analyzer = s.profiler.Push("[Req]" + rpcRequest.RpcRequestData.GetServiceMethod())
				}

				s.GetRpcHandler().HandlerRpcRequest(rpcRequest)
				if analyzer != nil {
					analyzer.Pop()
					analyzer = nil
				}
				event.DeleteEvent(cEvent)
			case event.ServiceRpcResponseEvent:
				cEvent, ok := ev.(*event.Event)
				if ok == false {
					log.Error("Type event conversion error")
					break
				}
				rpcResponseCB, ok := cEvent.Data.(*rpc.Call)
				if ok == false {
					log.Error("Type *rpc.Call conversion error")
					break
				}
				if s.profiler != nil {
					analyzer = s.profiler.Push("[Res]" + rpcResponseCB.ServiceMethod)
				}
				s.GetRpcHandler().HandlerRpcResponseCB(rpcResponseCB)
				if analyzer != nil {
					analyzer.Pop()
					analyzer = nil
				}
				event.DeleteEvent(cEvent)
			default:
				if s.profiler != nil {
					analyzer = s.profiler.Push("[SEvent]" + strconv.Itoa(int(ev.GetEventType())))
				}
				s.eventProcessor.EventHandler(ev)
				if analyzer != nil {
					analyzer.Pop()
					analyzer = nil
				}
			}

		case t := <-s.dispatcher.ChanTimer:
			if s.profiler != nil {
				analyzer = s.profiler.Push("[timer]" + t.GetName())
			}
			t.Do()
			if analyzer != nil {
				analyzer.Pop()
				analyzer = nil
			}
		}

		if bStop == true {
			if atomic.AddInt32(&s.goroutineNum, -1) <= 0 {
				s.startStatus = false
				s.Release()
			}
			break
		}
	}
}

/*	=====Implement IService=====	 */

func (s *Service) OnSetup(iService IService) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) OnInit() error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) OnStart() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) OnRelease() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SetName(serviceName string) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetServiceCfg() interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetProfiler() *interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetServiceEventChannelNum() int {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetServiceTimerChannelNum() int {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SetEventChannelNum(num int) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) OpenProfiler() {
	//TODO implement me
	panic("implement me")
}

// Release 关闭和释放资源
func (s *Service) Release() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			l := runtime.Stack(buf, false)
			errString := fmt.Sprint(r)
			log.Dump(string(buf[:l]), log.String("error", errString))
		}
	}()

	s.self.OnRelease()
}