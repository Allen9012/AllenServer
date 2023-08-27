package network

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 封装pack中的消息
  @modified by:
**/

type Message struct {
	ID   uint64
	Data []byte
}
