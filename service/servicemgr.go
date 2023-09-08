package service

import "errors"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: 服务注册
  @modified by:
**/

// 本地所有的service
var mapServiceName map[string]IService
var setupServiceList []IService

type RegRpcEventFunType func(serviceName string)
type RegDiscoveryServiceEventFunType func(serviceName string)

var RegRpcEventFun RegRpcEventFunType
var UnRegRpcEventFun RegRpcEventFunType

var RegDiscoveryServiceEventFun RegDiscoveryServiceEventFunType
var UnRegDiscoveryServiceEventFun RegDiscoveryServiceEventFunType

func init() {
	mapServiceName = map[string]IService{}
	setupServiceList = []IService{}
}

// Init 初始化每一个service
func Init() {
	for _, s := range setupServiceList {
		err := s.OnInit()
		if err != nil {
			errs := errors.New("Failed to initialize " + s.GetName() + " service:" + err.Error())
			panic(errs)
		}
	}
}

// Setup 注册一个service服务
func Setup(s IService) bool {
	_, ok := mapServiceName[s.GetName()]
	if ok == true {
		return false
	}

	mapServiceName[s.GetName()] = s
	setupServiceList = append(setupServiceList, s)
	return true
}

func GetService(serviceName string) IService {
	s, ok := mapServiceName[serviceName]
	if ok == false {
		return nil
	}

	return s
}

func Start() {
	for _, s := range setupServiceList {
		s.Start()
	}
}

func StopAllService() {
	for i := len(setupServiceList) - 1; i >= 0; i-- {
		setupServiceList[i].Stop()
	}
}
