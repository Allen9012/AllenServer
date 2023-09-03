package friend

import (
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"google.golang.org/protobuf/proto"
)

type Owner interface {
	Start()
	Stop()
	SendMsg(ID messageID.MessageId, message proto.Message)
}
