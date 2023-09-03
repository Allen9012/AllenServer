package task

import (
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"google.golang.org/protobuf/proto"
)

type Player interface {
	SendMsg(ID messageID.MessageId, message proto.Message)
	GetTaskData() *Data
}
