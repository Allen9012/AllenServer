package simple_event

import "github.com/Allen9012/AllenGame/event"

// 事件是origin中一个重要的组成部分，可以在同一个node中的service与service或者与module之间进行事件通知。
// 系统内置的几个服务，如：TcpService/HttpService等都是通过事件功能实现。
// 他也是一个典型的观察者设计模型。
// 在event中有两个类型的interface，一个是event.IEventProcessor它提供注册与卸载功能，另一个是event.IEventHandler提供消息广播等功能。

const (
	//自定义事件类型，必需从event.Sys_Event_User_Define开始

	//自定义事件1
	SimpleEvent1 = event.Sys_Event_User_Define + 1
)
