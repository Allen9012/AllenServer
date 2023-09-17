package simple_service

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: 间隔执行和定时任务
  @modified by:
**/

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
	"github.com/Allen9012/AllenGame/util/timer"
	"time"
)

// 模块加载时自动安装TestService2服务
func init() {
	node.Setup(&TestService2{})
}

// 新建自定义服务TestService1
type TestService2 struct {
	//所有的自定义服务必需加入service.Service基服务
	//那么该自定义服务将有各种功能特性
	//例如: Rpc,事件驱动,定时器等
	service.Service
}

// 服务初始化函数，在安装服务时，服务将自动调用OnInit函数
func (slf *TestService2) OnInit() error {
	fmt.Printf("【间隔执行和定时任务】启动\n")

	//间隔执行
	slf.AfterFunc(time.Second*1, slf.OnSecondTick)

	//crontab模式定时触发
	//NewCronExpr的参数分别代表:Seconds Minutes Hours DayOfMonth Month DayOfWeek
	//以下为每换分钟时触发
	cron, _ := timer.NewCronExpr("0 * * * * *")
	slf.CronFunc(cron, slf.OnCron)

	return nil
}

// OnSecondTick 间隔执行
func (slf *TestService2) OnSecondTick(t *timer.Timer) {
	fmt.Printf("tick.\n")
	slf.AfterFunc(time.Second*1, slf.OnSecondTick)
}

// OnCron 定时执行
func (slf *TestService2) OnCron(cron *timer.Cron) {
	fmt.Printf("A minute passed!\n")
	//cron.Close()
}
