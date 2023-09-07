package friend

import (
	"github.com/Allen9012/AllenServer/internal/event"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"google.golang.org/protobuf/proto"
)

type IPlayer interface {
	Start()
	Stop()
	SendMsg(ID messageID.MessageId, message proto.Message)
	OnEvent(event event.IEvent)
}
