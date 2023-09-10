package rpc

import (
	"github.com/Allen9012/AllenGame/util/sync"
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

type RpcRequest struct {
	ref            bool
	RpcRequestData IRpcRequestData

	inParam    interface{}
	localReply interface{}

	requestHandle RequestHandler
	callback      *reflect.Value
	rpcProcessor  IRpcProcessor
}

type RpcResponse struct {
	RpcResponseData IRpcResponseData
}

type Responder = RequestHandler

func (r *Responder) IsInvalid() bool {
	return reflect.ValueOf(*r).Pointer() == reflect.ValueOf(reqHandlerNull).Pointer()
}

var rpcRequestPool = sync.NewPoolEx(make(chan sync.IPoolData, 10240), func() sync.IPoolData {
	return &RpcRequest{}
})

var rpcCallPool = sync.NewPoolEx(make(chan sync.IPoolData, 10240), func() sync.IPoolData {
	return &Call{done: make(chan *Call, 1)}
})

type IRpcRequestData interface {
	GetSeq() uint64
	GetServiceMethod() string
	GetInParam() []byte
	IsNoReply() bool
	GetRpcMethodId() uint32
}

type IRpcResponseData interface {
	GetSeq() uint64
	GetErr() *RpcError
	GetReply() []byte
}

type RpcHandleFinder interface {
	FindRpcHandler(serviceMethod string) IRpcHandler
}

type RequestHandler func(Returns interface{}, Err RpcError)

type Call struct {
	ref           bool
	Seq           uint64
	ServiceMethod string
	Reply         interface{}
	Response      *RpcResponse
	Err           error
	done          chan *Call // Strobes when call is complete.
	connId        int
	callback      *reflect.Value
	rpcHandler    IRpcHandler
	TimeOut       time.Duration
}

type RpcCancel struct {
	Cli     *Client
	CallSeq uint64
}

/*	RpcRequest Implement sync.IPoolData */

func (r *RpcRequest) Reset() {
	//TODO implement me
	panic("implement me")
}

func (r *RpcRequest) IsRef() bool {
	//TODO implement me
	panic("implement me")
}

func (r *RpcRequest) Ref() {
	//TODO implement me
	panic("implement me")
}

func (r *RpcRequest) UnRef() {
	//TODO implement me
	panic("implement me")
}

/*	Call Implement sync.IPoolData */

func (c *Call) Reset() {
	//TODO implement me
	panic("implement me")
}

func (c *Call) IsRef() bool {
	//TODO implement me
	panic("implement me")
}

func (c *Call) Ref() {
	//TODO implement me
	panic("implement me")
}

func (c *Call) UnRef() {
	//TODO implement me
	panic("implement me")
}
