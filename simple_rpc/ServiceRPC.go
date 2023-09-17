package simple_rpc

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
	"github.com/Allen9012/AllenGame/util/timer"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/17
  @desc:
  @modified by:
**/

func init() {
	node.Setup(&ServiceRPC{}, &ServiceRPCTest{})
}

type ServiceRPC struct {
	service.Service
}

//----------------------------------------------------------------------------------------------

type ServiceRPCTest struct {
	service.Service
}

func (slf *ServiceRPCTest) OnInit() error {
	fmt.Printf("【RPC测试服务】启动\n")
	slf.AfterFunc(time.Second*2, slf.CallRPC)
	slf.AfterFunc(time.Second*6, slf.AsyncCallRPC)
	slf.AfterFunc(time.Second*9, slf.GoCall)
	return nil
}

// 同步rpc
func (slf *ServiceRPCTest) CallRPC(t *timer.Timer) {
	fmt.Printf("CallRPC--------------------\n")

	var input InputData
	input.A = 10
	input.B = 22

	var output int

	//1，同步方式

	//同步调用其他服务的rpc,input为传入的rpc,output为输出参数
	err := slf.Call("ServiceRPC.RPCSum", &input, &output)
	if err != nil {
		fmt.Printf("Call-err: %v\n", err)
	} else {
		fmt.Printf("Call-output: %v\n", output)
	}

	//自定义超时,默认rpc超时时间为15s
	err = slf.CallWithTimeout(time.Second*1, "ServiceRPC.RPCSum", &input, &output)
	if err != nil {
		fmt.Printf("CallWithTimeout-err: %v\n", err)
	} else {
		fmt.Printf("CallWithTimeout-output: %v\n", output)
	}
}

// 异步rpc
func (slf *ServiceRPCTest) AsyncCallRPC(t *timer.Timer) {
	fmt.Printf("AsyncCallRPC--------------------\n")
	var input InputData
	input.A = 300
	input.B = 600
	/*slf.AsyncCallNode(1,"TestService6.RPC_Sum",&input,func(output *int,err error){
	})*/
	//异步调用，在数据返回时，会回调传入函数
	//注意函数的第一个参数一定是RPC_Sum函数的第二个参数，err error为RPC_Sum返回值
	err := slf.AsyncCall("ServiceRPC.RPCSum", &input, func(output *int, err error) {
		if err != nil {
			fmt.Printf("AsyncCall error :%+v\n", err)
		} else {
			fmt.Printf("AsyncCall output %d\n", *output)
		}
	})
	if err != nil {
		fmt.Println(err)
	}

	//自定义超时,返回一个cancel函数，可以在业务需要时取消rpc调用
	rpcCancel, err := slf.AsyncCallWithTimeout(time.Second*1, "ServiceRPC.RPCSum", &input, func(output *int, err error) {
		//如果下面注释的rpcCancel()函数被调用，这里可能将不再返回
		if err != nil {
			fmt.Printf("AsyncCallWithTimeout error :%+v\n", err)
		} else {
			fmt.Printf("AsyncCallWithTimeout output %d\n", *output)
		}
	})

	if err != nil {
		fmt.Println(err)

		if rpcCancel != nil {
			fmt.Printf("rpcCancel: %v\n", rpcCancel)
		}
	}
}

// rpc广播
func (slf *ServiceRPCTest) GoCall(t *timer.Timer) {
	fmt.Printf("GoCall--------------------\n")
	var input InputData
	input.A = 30120
	input.B = 6100

	//在某些应用场景下不需要数据返回可以使用Go，它是不阻塞的,只需要填入输入参数
	err := slf.Go("ServiceRPC.RPCSum", &input)
	if err != nil {
		fmt.Printf("Go error :%+v\n", err)
	}

	//以下是广播方式，如果在同一个子网中有多个同名的服务名，CastGo将会广播给所有的node
	slf.CastGo("ServiceRPC.RPCSum", &input)
}
