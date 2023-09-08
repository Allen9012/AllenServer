package rpc

import (
	"reflect"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

type IRpcHandlerChannel interface {
	PushRpcResponse(call *Call) error
	PushRpcRequest(rpcRequest *RpcRequest) error
}

type IRpcHandler interface {
	IRpcHandlerChannel
	GetName() string
	InitRpcHandler(rpcHandler IRpcHandler, getClientFun FuncRpcClient, getServerFun FuncRpcServer, rpcHandlerChannel IRpcHandlerChannel)
	GetRpcHandler() IRpcHandler
	HandlerRpcRequest(request *RpcRequest)
	HandlerRpcResponseCB(call *Call)
	CallMethod(client *Client, ServiceMethod string, param interface{}, callBack reflect.Value, reply interface{}) error

	Call(serviceMethod string, args interface{}, reply interface{}) error
	CallNode(nodeId int, serviceMethod string, args interface{}, reply interface{}) error
	AsyncCall(serviceMethod string, args interface{}, callback interface{}) error
	AsyncCallNode(nodeId int, serviceMethod string, args interface{}, callback interface{}) error

	CallWithTimeout(timeout time.Duration, serviceMethod string, args interface{}, reply interface{}) error
	CallNodeWithTimeout(timeout time.Duration, nodeId int, serviceMethod string, args interface{}, reply interface{}) error
	AsyncCallWithTimeout(timeout time.Duration, serviceMethod string, args interface{}, callback interface{}) (CancelRpc, error)
	AsyncCallNodeWithTimeout(timeout time.Duration, nodeId int, serviceMethod string, args interface{}, callback interface{}) (CancelRpc, error)

	Go(serviceMethod string, args interface{}) error
	GoNode(nodeId int, serviceMethod string, args interface{}) error
	RawGoNode(rpcProcessorType RpcProcessorType, nodeId int, rpcMethodId uint32, serviceName string, rawArgs []byte) error
	CastGo(serviceMethod string, args interface{}) error
	IsSingleCoroutine() bool
	UnmarshalInParam(rpcProcessor IRpcProcessor, serviceMethod string, rawRpcMethodId uint32, inParam []byte) (interface{}, error)
	GetRpcServer() FuncRpcServer
}
