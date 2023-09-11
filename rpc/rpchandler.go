package rpc

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: rpc handler
  @modified by:
**/

const maxClusterNode int = 128

type FuncRpcClient func(nodeId int, serviceMethod string, client []*Client) (error, int)
type FuncRpcServer func() *Server

/*	定义RPC Error相关	Implement error */
var nilError = reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())

type RpcError string

var NilError RpcError

func (e RpcError) Error() string {
	return string(e)
}

func ConvertError(e error) RpcError {
	if e == nil {
		return NilError
	}

	rpcErr := RpcError(e.Error())
	return rpcErr
}

type RpcMethodInfo struct {
	method           reflect.Method
	inParamValue     reflect.Value
	inParam          interface{}
	outParamValue    reflect.Value
	hasResponder     bool
	rpcProcessorType RpcProcessorType
}

type RawRpcCallBack func(rawData []byte)

type IRpcHandlerChannel interface {
	PushRpcResponse(call *Call) error
	PushRpcRequest(rpcRequest *RpcRequest) error
}

type IDiscoveryServiceListener interface {
	OnDiscoveryService(nodeId int, serviceName []string)
	OnUnDiscoveryService(nodeId int, serviceName []string)
}

type CancelRpc func()

func emptyCancelRpc() {}

type IRpcHandler interface {
	IRpcHandlerChannel
	GetName() string
	InitRpcHandler(rpcHandler IRpcHandler, getClientFun FuncRpcClient, getServerFun FuncRpcServer, rpcHandlerChannel IRpcHandlerChannel)
	GetRpcHandler() IRpcHandler
	HandlerRpcRequest(request *RpcRequest)
	HandlerRpcResponseCB(call *Call)
	CallMethod(client *Client, ServiceMethod string, param interface{}, callBack reflect.Value, reply interface{}) error
	// 同步等待调用结果
	Call(serviceMethod string, args interface{}, reply interface{}) error
	CallNode(nodeId int, serviceMethod string, args interface{}, reply interface{}) error
	// 异步, rpc首选, 不会阻塞本服务
	AsyncCall(serviceMethod string, args interface{}, callback interface{}) error
	// 在明确节点时调用,可以稍微减少开销;  Service名相同时, 避免广播
	AsyncCallNode(nodeId int, serviceMethod string, args interface{}, callback interface{}) error

	CallWithTimeout(timeout time.Duration, serviceMethod string, args interface{}, reply interface{}) error
	CallNodeWithTimeout(timeout time.Duration, nodeId int, serviceMethod string, args interface{}, reply interface{}) error
	AsyncCallWithTimeout(timeout time.Duration, serviceMethod string, args interface{}, callback interface{}) (CancelRpc, error)
	AsyncCallNodeWithTimeout(timeout time.Duration, nodeId int, serviceMethod string, args interface{}, callback interface{}) (CancelRpc, error)
	// 无结果,不阻塞
	Go(serviceMethod string, args interface{}) error
	GoNode(nodeId int, serviceMethod string, args interface{}) error
	// 原数据,减少参数/结果的序列化和反序列化, 大量转发时使用.
	RawGoNode(rpcProcessorType RpcProcessorType, nodeId int, rpcMethodId uint32, serviceName string, rawArgs []byte) error
	// 广播
	CastGo(serviceMethod string, args interface{}) error
	IsSingleCoroutine() bool
	UnmarshalInParam(rpcProcessor IRpcProcessor, serviceMethod string, rawRpcMethodId uint32, inParam []byte) (interface{}, error)
	GetRpcServer() FuncRpcServer
}

/*	Implement IRpcHandler */
type RpcHandler struct {
	IRpcHandlerChannel

	rpcHandler      IRpcHandler
	mapFunctions    map[string]RpcMethodInfo
	mapRawFunctions map[uint32]RawRpcCallBack
	funcRpcClient   FuncRpcClient
	funcRpcServer   FuncRpcServer

	pClientList []*Client
}

// INodeListener 定义节点监听接口
type INodeListener interface {
	OnNodeConnected(nodeId int)
	OnNodeDisconnect(nodeId int)
}

func reqHandlerNull(Returns interface{}, Err RpcError) {
}

var requestHandlerNull reflect.Value

func init() {
	requestHandlerNull = reflect.ValueOf(reqHandlerNull)
}

/*	=====Implement IRpcHandler=====*/

func (handler *RpcHandler) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) InitRpcHandler(rpcHandler IRpcHandler, getClientFun FuncRpcClient, getServerFun FuncRpcServer, rpcHandlerChannel IRpcHandlerChannel) {
	handler.IRpcHandlerChannel = rpcHandlerChannel
	handler.mapRawFunctions = make(map[uint32]RawRpcCallBack)
	handler.rpcHandler = rpcHandler
	handler.mapFunctions = map[string]RpcMethodInfo{}
	handler.funcRpcClient = getClientFun
	handler.funcRpcServer = getServerFun
	handler.pClientList = make([]*Client, maxClusterNode)
	handler.RegisterRpc(rpcHandler)
}

func (handler *RpcHandler) GetRpcHandler() IRpcHandler {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) HandlerRpcRequest(request *RpcRequest) {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) HandlerRpcResponseCB(call *Call) {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) CallMethod(client *Client, ServiceMethod string, param interface{}, callBack reflect.Value, reply interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) Call(serviceMethod string, args interface{}, reply interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) CallNode(nodeId int, serviceMethod string, args interface{}, reply interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) AsyncCall(serviceMethod string, args interface{}, callback interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) AsyncCallNode(nodeId int, serviceMethod string, args interface{}, callback interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) CallWithTimeout(timeout time.Duration, serviceMethod string, args interface{}, reply interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) CallNodeWithTimeout(timeout time.Duration, nodeId int, serviceMethod string, args interface{}, reply interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) AsyncCallWithTimeout(timeout time.Duration, serviceMethod string, args interface{}, callback interface{}) (CancelRpc, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) AsyncCallNodeWithTimeout(timeout time.Duration, nodeId int, serviceMethod string, args interface{}, callback interface{}) (CancelRpc, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) Go(serviceMethod string, args interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) GoNode(nodeId int, serviceMethod string, args interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) RawGoNode(rpcProcessorType RpcProcessorType, nodeId int, rpcMethodId uint32, serviceName string, rawArgs []byte) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) CastGo(serviceMethod string, args interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) IsSingleCoroutine() bool {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) UnmarshalInParam(rpcProcessor IRpcProcessor, serviceMethod string, rawRpcMethodId uint32, inParam []byte) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) GetRpcServer() FuncRpcServer {
	//TODO implement me
	panic("implement me")
}

func (handler *RpcHandler) RegisterRpc(rpcHandler IRpcHandler) error {
	typ := reflect.TypeOf(rpcHandler)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		err := handler.suitableMethods(method)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (handler *RpcHandler) suitableMethods(method reflect.Method) error {
	//只有RPC_开头的才能被调用
	if strings.Index(method.Name, "RPC_") != 0 && strings.Index(method.Name, "RPC") != 0 {
		return nil
	}

	//取出输入参数类型
	var rpcMethodInfo RpcMethodInfo
	typ := method.Type
	if typ.NumOut() != 1 {
		return fmt.Errorf("%s The number of returned arguments must be 1", method.Name)
	}

	if typ.Out(0).String() != "error" {
		return fmt.Errorf("%s The return parameter must be of type error", method.Name)
	}

	if typ.NumIn() < 2 || typ.NumIn() > 4 {
		return fmt.Errorf("%s Unsupported parameter format", method.Name)
	}

	//1.判断第一个参数
	var parIdx = 1
	if typ.In(parIdx).String() == "rpc.RequestHandler" {
		parIdx += 1
		rpcMethodInfo.hasResponder = true
	}

	for i := parIdx; i < typ.NumIn(); i++ {
		if handler.isExportedOrBuiltinType(typ.In(i)) == false {
			return fmt.Errorf("%s Unsupported parameter types", method.Name)
		}
	}

	rpcMethodInfo.inParamValue = reflect.New(typ.In(parIdx).Elem())
	rpcMethodInfo.inParam = reflect.New(typ.In(parIdx).Elem()).Interface()
	pt, _ := GetProcessorType(rpcMethodInfo.inParamValue.Interface())
	rpcMethodInfo.rpcProcessorType = pt

	parIdx++
	if parIdx < typ.NumIn() {
		rpcMethodInfo.outParamValue = reflect.New(typ.In(parIdx).Elem())
	}

	rpcMethodInfo.method = method
	handler.mapFunctions[handler.rpcHandler.GetName()+"."+method.Name] = rpcMethodInfo
	return nil
}

// Is this an exported - upper case - name?
func isExported(name string) bool {
	r, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(r)
}

// Is this type exported or a builtin?
func (handler *RpcHandler) isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}
