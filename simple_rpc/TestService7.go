package simple_rpc

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/11
  @desc:
  @modified by:
**/

func init() {
	node.Setup(&TestService7{})
}

type TestService7 struct {
	service.Service
}

func (slf *TestService7) OnInit() error {
	slf.AfterFunc(time.Second*2, slf.CallTest)
	slf.AfterFunc(time.Second*2, slf.AsyncCallTest)
	slf.AfterFunc(time.Second*2, slf.GoTest)
	return nil
}

func (slf *TestService7) CallTest() {
	var input InputData
	input.A = 300
	input.B = 600
	var output int

	//同步调用其他服务的rpc,input为传入的rpc,output为输出参数
	err := slf.Call("TestService6.RPC_Sum", &input, &output)
	if err != nil {
		fmt.Printf("Call error :%+v\n", err)
	} else {
		fmt.Printf("Call output %d\n", output)
	}

	//自定义超时,默认rpc超时时间为15s
	err = slf.CallWithTimeout(time.Second*1, "TestService6.RPC_Sum", &input, &output)
	if err != nil {
		fmt.Printf("Call error :%+v\n", err)
	} else {
		fmt.Printf("Call output %d\n", output)
	}
}

func (slf *TestService7) AsyncCallTest() {
	var input InputData
	input.A = 300
	input.B = 600
	/*slf.AsyncCallNode(1,"TestService6.RPC_Sum",&input,func(output *int,err error){
	})*/
	//异步调用，在数据返回时，会回调传入函数
	//注意函数的第一个参数一定是RPC_Sum函数的第二个参数，err error为RPC_Sum返回值
	err := slf.AsyncCall("TestService6.RPC_Sum", &input, func(output *int, err error) {
		if err != nil {
			fmt.Printf("AsyncCall error :%+v\n", err)
		} else {
			fmt.Printf("AsyncCall output %d\n", *output)
		}
	})
	fmt.Println(err)

	//自定义超时,返回一个cancel函数，可以在业务需要时取消rpc调用
	rpcCancel, err := slf.AsyncCallWithTimeout(time.Second*1, "TestService6.RPC_Sum", &input, func(output *int, err error) {
		//如果下面注释的rpcCancel()函数被调用，这里可能将不再返回
		if err != nil {
			fmt.Printf("AsyncCall error :%+v\n", err)
		} else {
			fmt.Printf("AsyncCall output %d\n", *output)
		}
	})
	//rpcCancel()
	fmt.Println(err, rpcCancel)

}

func (slf *TestService7) GoTest() {
	var input InputData
	input.A = 300
	input.B = 600

	//在某些应用场景下不需要数据返回可以使用Go，它是不阻塞的,只需要填入输入参数
	err := slf.Go("TestService6.RPC_Sum", &input)
	if err != nil {
		fmt.Printf("Go error :%+v\n", err)
	}

	//以下是广播方式，如果在同一个子网中有多个同名的服务名，CastGo将会广播给所有的node
	//slf.CastGo("TestService6.RPC_Sum",&input)
}
