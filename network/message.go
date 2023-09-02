package network

import "sync"

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

var msgPool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}

// GetPooledMessage gets a pooled Message.
func GetPooledMessage() *Message {
	return msgPool.Get().(*Message)
}

// FreeMessage puts a Message into the pool.
func FreeMessage(msg *Message) {
	if msg != nil && len(msg.Data) > 0 {
		ResetMessage(msg)
		msgPool.Put(msg)
	}
}

// ResetMessage reset a Message
func ResetMessage(m *Message) {
	m.Data = m.Data[:0]
}
