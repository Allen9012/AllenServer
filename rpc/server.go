package rpc

import "github.com/Allen9012/AllenGame/network"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

type RpcProcessorType uint8

const (
	RpcProcessorJson   RpcProcessorType = 0
	RpcProcessorGoGoPB RpcProcessorType = 1
)

// var arrayProcessor = []IRpcProcessor{&JsonProcessor{}, &PBProcessor{}}
var arrayProcessorLen uint8 = 2
var LittleEndian bool

type Server struct {
	functions       map[interface{}]interface{}
	rpcHandleFinder RpcHandleFinder
	rpcServer       *network.TCPServer

	compressBytesLen int
}

func (server *Server) Init(rpcHandleFinder RpcHandleFinder) {
	server.rpcHandleFinder = rpcHandleFinder
	server.rpcServer = &network.TCPServer{}
}
