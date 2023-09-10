package rpc

import "github.com/Allen9012/AllenGame/network"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

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
