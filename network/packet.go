package network

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/30
  @desc:
  @modified by:
**/

type Packet struct {
	Msg  *Message
	Conn *TcpConnX
}
