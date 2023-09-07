package chat

import (
	"errors"
	"fmt"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"github.com/Allen9012/AllenServer/network/protocol/gen/player"
	"google.golang.org/protobuf/proto"
	"sync"
)

type PrivateChatHandler struct {
	Id messageID.MessageId
	Fn func(p *PrivateChat, packet *network.Message)
}

var (
	handlers     []*PrivateChatHandler
	onceInit     sync.Once
	MinMessageId messageID.MessageId
	MaxMessageId messageID.MessageId
)

func init() {
	onceInit.Do(func() {
		HandlerChatRegister()
	})
}

func GetHandler(id messageID.MessageId) (*PrivateChatHandler, error) {
	if id > MinMessageId && id < MaxMessageId {
		return nil, errors.New("not in")
	}
	for _, handler := range handlers {
		if handler.Id == id {
			return handler, nil
		}
	}
	return nil, errors.New("not exist")
}

func HandlerChatRegister() {
	handlers[0] = &PrivateChatHandler{
		Id: messageID.MessageId_SCSendChatMsg,
		Fn: ResolvePrivateChatMsg,
	}
}

func ResolvePrivateChatMsg(p *PrivateChat, packet *network.Message) {
	req := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	fmt.Println(req.Msg.Content)
	p.SendMsg(messageID.MessageId_SCSendChatMsg, &player.SCSendChatMsg{})
}
