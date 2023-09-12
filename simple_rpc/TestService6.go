package simple_rpc

import (
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

func (slf *TestService6) OnInit() error {
	return nil
}

type InputData struct {
	A int
	B int
}

// 注意RPC函数名的格式必需为RPC_FunctionName或者是RPCFunctionName，如下的RPC_Sum也可以写成RPCSum
func (slf *TestService6) RPC_Sum(input *InputData, output *int) error {
	*output = input.A + input.B
	return nil
}
