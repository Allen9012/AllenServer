package network

import "net"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/30
  @desc:
  @modified by:
**/

type ClientPacket struct {
	Msg  *Message
	Conn net.Conn
}

type SessionPacket struct {
	Msg  *Message
	Sess *Session
}