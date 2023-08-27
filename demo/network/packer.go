package network

import (
	"encoding/binary"
	"io"
	"net"
	"sync"
	"time"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc: 定义pack的形式收发数据
	@modified by:
*/

var (
	once                 sync.Once
	NormalPackerInstance *NormalPacker
)

type NormalPacker struct {
	Order binary.ByteOrder // 大端或者小端存储顺序
}

const LEN_UINT64 = 8

// NewNormalPacker 增加了修改我们使得Packer需要成为单例
func init() {
	once.Do(func() {
		NormalPackerInstance = &NormalPacker{
			Order: binary.BigEndian,
		}
	})
}

// Pack    |data 长度 | id | data|
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	// 给定了长度的buffer
	buffer := make([]byte, LEN_UINT64+LEN_UINT64+len(message.Data))
	p.Order.PutUint64(buffer[:8], uint64(len(message.Data)))
	p.Order.PutUint64(buffer[8:16], message.ID)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

func (p *NormalPacker) Unpack(reader io.Reader) (*Message, error) {
	// 设置超时时间
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 8+8)
	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}
	TotalLen := p.Order.Uint64(buffer[:8])
	id := p.Order.Uint64(buffer[8:16])
	DataLen := TotalLen - 16
	DataBuffer := make([]byte, DataLen)
	_, err = io.ReadFull(reader, DataBuffer)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		ID:   id,
		Data: DataBuffer,
	}
	return msg, nil
}
