package rpc

import (
	"github.com/Allen9012/AllenGame/util/sync"
	jsoniter "github.com/json-iterator/go"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/11
  @desc:
  @modified by:
**/

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JsonProcessor struct {
}

type JsonRpcRequestData struct {
	//packhead
	Seq           uint64 // sequence number chosen by client
	rpcMethodId   uint32
	ServiceMethod string // format: "Service.Method"
	NoReply       bool   //是否需要返回
	//packbody
	InParam []byte
}

type JsonRpcResponseData struct {
	//head
	Seq uint64 // sequence number chosen by client
	Err string

	//returns
	Reply []byte
}

/*	池化技术	*/

var rpcJsonResponseDataPool = sync.NewPool(make(chan interface{}, 10240), func() interface{} {
	return &JsonRpcResponseData{}
})

var rpcJsonRequestDataPool = sync.NewPool(make(chan interface{}, 10240), func() interface{} {
	return &JsonRpcRequestData{}
})

/*	Implement IRpcProcessor	 */

func (jsonProcessor *JsonProcessor) Clone(src interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) Marshal(v interface{}) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) Unmarshal(data []byte, v interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) MakeRpcRequest(seq uint64, rpcMethodId uint32, serviceMethod string, noReply bool, inParam []byte) IRpcRequestData {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) MakeRpcResponse(seq uint64, err RpcError, reply []byte) IRpcResponseData {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) ReleaseRpcRequest(rpcRequestData IRpcRequestData) {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) ReleaseRpcResponse(rpcRequestData IRpcResponseData) {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) IsParse(param interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (jsonProcessor *JsonProcessor) GetProcessorType() RpcProcessorType {
	//TODO implement me
	panic("implement me")
}
