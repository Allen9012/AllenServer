package task

import (
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type Status int

const (
	ACCEPT Status = iota + 1
	ING
	FINISH
	SUBMIT
)

const (
	defaultLoopNum     = 50
	defaultMonitor     = 100
	defaultChanInSize  = 1000
	defaultChanOutSize = 500
)

type ManagerConfig struct {
	LoopNum    int
	MonitorNum int
	ChInSize   int
	ChOutSize  int
}

type Config struct {
	Id              uint32        `json:"id"`
	Name            string        `json:"name"`
	DropId          uint32        `json:"dropId"` //掉落
	Category        int           `json:"category"`
	Targets         []*TargetConf `json:"targets"`
	SubmitType      int           `json:"submitType"` //自动提交，手动提交
	AcceptType      int           `json:"acceptType"`
	CompleteNtf     int           `json:"completeNtf"` //完成是否推送
	UnlockCondition int           `json:"unlockCondition"`
}

type TargetConf struct {
	Id            uint32
	DropId        uint32
	Name          string
	CompleteParam string
}

type PlayerActionParam struct {
	MessageId messageID.MessageId
	Player    Player
	Packet    *network.Message
}

type EventParam struct {
	EventCategory
	Player Player
	Param  interface{}
}

type TargetCategory int

const (
	TargetCategoryNormal TargetCategory = iota + 1
)

type EventCategory int

const (
	NormalEvent EventCategory = iota + 1
	BaseEvent
)
