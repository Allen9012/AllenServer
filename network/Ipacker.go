package network

import "io"

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:
*/

type IPacker interface {
	Pack(message *Message) ([]byte, error)
	Read(*TcpSession) ([]byte, error)
	Unpack(reader io.Reader) (*Message, error)
}
