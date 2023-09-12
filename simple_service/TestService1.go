package simple_service

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

// 模块加载时自动安装TestService1服务
func init() {
	node.Setup(&TestService1{})
}

// 新建自定义服务TestService1
type TestService1 struct {

	//所有的自定义服务必需加入service.Service基服务
	//那么该自定义服务将有各种功能特性
	//例如: Rpc,事件驱动,定时器等
	service.Service
}

// 服务初始化函数，在安装服务时，服务将自动调用OnInit函数
func (slf *TestService1) OnInit() error {
	fmt.Printf("TestService1 OnInit.\n")

	//打开多线程处理模式，10个协程并发处理
	slf.SetGoRoutineNum(10)
	return nil
}
