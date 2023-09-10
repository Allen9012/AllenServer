package rpc

import (
	"github.com/Allen9012/AllenGame/network"
	"reflect"
	"sync"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

type IRealClient interface {
	SetConn(conn *network.TCPConn)
	Close(waitDone bool)

	AsyncCall(timeout time.Duration, rpcHandler IRpcHandler, serviceMethod string, callback reflect.Value, args interface{}, replyParam interface{}, cancelable bool) (CancelRpc, error)
	Go(timeout time.Duration, rpcHandler IRpcHandler, noReply bool, serviceMethod string, args interface{}, reply interface{}) *Call
	RawGo(timeout time.Duration, rpcHandler IRpcHandler, processor IRpcProcessor, noReply bool, rpcMethodId uint32, serviceMethod string, rawArgs []byte, reply interface{}) *Call
	IsConnected() bool

	Run()
	OnClose()
}

type Client struct {
	clientId             uint32
	nodeId               int
	pendingLock          sync.RWMutex
	startSeq             uint64
	pending              map[uint64]*Call
	callRpcTimeout       time.Duration
	maxCheckCallRpcCount int

	callTimerHeap CallTimerHeap
	IRealClient   // 组合
}

func (c *Client) SetConn(conn *network.TCPConn) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Close(waitDone bool) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) AsyncCall(timeout time.Duration, rpcHandler IRpcHandler, serviceMethod string, callback reflect.Value, args interface{}, replyParam interface{}, cancelable bool) (CancelRpc, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Go(timeout time.Duration, rpcHandler IRpcHandler, noReply bool, serviceMethod string, args interface{}, reply interface{}) *Call {
	//TODO implement me
	panic("implement me")
}

func (c *Client) RawGo(timeout time.Duration, rpcHandler IRpcHandler, processor IRpcProcessor, noReply bool, rpcMethodId uint32, serviceMethod string, rawArgs []byte, reply interface{}) *Call {
	//TODO implement me
	panic("implement me")
}

func (c *Client) IsConnected() bool {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Run() {
	//TODO implement me
	panic("implement me")
}

func (c *Client) OnClose() {
	//TODO implement me
	panic("implement me")
}
