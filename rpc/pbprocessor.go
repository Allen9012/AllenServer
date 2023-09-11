package rpc

import "github.com/Allen9012/AllenGame/util/sync"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/11
  @desc:
  @modified by:
**/

type PBProcessor struct {
}

/*	池化技术	*/

var rpcPbResponseDataPool = sync.NewPool(make(chan interface{}, 10240), func() interface{} {
	return &PBRpcResponseData{}
})

var rpcPbRequestDataPool = sync.NewPool(make(chan interface{}, 10240), func() interface{} {
	return &PBRpcRequestData{}
})

/*	Implement IRpcProcessor	 */

func (P *PBProcessor) Clone(src interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) Marshal(v interface{}) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) Unmarshal(data []byte, v interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) MakeRpcRequest(seq uint64, rpcMethodId uint32, serviceMethod string, noReply bool, inParam []byte) IRpcRequestData {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) MakeRpcResponse(seq uint64, err RpcError, reply []byte) IRpcResponseData {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) ReleaseRpcRequest(rpcRequestData IRpcRequestData) {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) ReleaseRpcResponse(rpcRequestData IRpcResponseData) {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) IsParse(param interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) GetProcessorType() RpcProcessorType {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) MsgRoute(clientId uint64, msg interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) UnknownMsgRoute(clientId uint64, msg interface{}) {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) ConnectedRoute(clientId uint64) {
	//TODO implement me
	panic("implement me")
}

func (P *PBProcessor) DisConnectedRoute(clientId uint64) {
	//TODO implement me
	panic("implement me")
}
