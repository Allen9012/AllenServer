package simple_rpc

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/11
  @desc:
  @modified by:
**/

func init() {
	node.Setup(&TestService6{})
}

type TestService6 struct {
	service.Service
}

//----------------------------------------------------------------------------------------------

type InputData struct {
	A, B int
}

// ----------------------------------------------------------------------------------------------
func (slf *TestService6) OnInit() error {
	fmt.Printf("【RPC服务】启动\n")
	return nil
}

// 注意RPC函数名的格式必需为RPC_FunctionName或者是RPCFunctionName，如下的RPC_Sum也可以写成RPCSum
func (slf *TestService6) RPCSum(input *InputData, output *int) error {
	fmt.Printf("ServiceRPC-RPCSum\n")
	*output = input.A + input.B
	return nil
}
