package network

import (
	"net"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: Conn
  @modified by:
**/

type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	ReleaseReadMsg(byteBuff []byte)
}
