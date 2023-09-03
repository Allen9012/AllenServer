package gateway

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type Server struct {
	cliServer *ClientServer
	inServer  *InnerServer
}

func NewServer() *Server {
	return &Server{
		cliServer: NewClientServer(),
		inServer:  NewInnerServer(),
	}
}

func (s *Server) Start() {

}
