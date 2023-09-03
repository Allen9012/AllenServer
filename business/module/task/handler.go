package task

import (
	"errors"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"sync"
)

type Handler struct {
	Id messageID.MessageId
	Fn func(player Player, packet *network.Message)
}

var (
	handlers     []*Handler
	onceInit     sync.Once
	MinMessageId messageID.MessageId
	MaxMessageId messageID.MessageId //handle 的消息范围
)

func IsBelongToHere(id messageID.MessageId) bool {
	return id > MinMessageId && id < MaxMessageId
}

func GetHandler(id messageID.MessageId) (*Handler, error) {
	for _, handler := range handlers {
		if handler.Id == id {
			return handler, nil
		}
	}
	return nil, errors.New("not exist")
}

func init() {
	onceInit.Do(func() {
		HandlerFriendRegister()
	})
}

func HandlerFriendRegister() {
	handlers[0] = &Handler{
		0,
		AcceptTask,
	}
	handlers[1] = &Handler{
		0,
		Submit,
	}
}

// AcceptTask accept task
func AcceptTask(player Player, packet *network.Message) {
	player.GetTaskData().GetTask(0).SetStatus(ACCEPT)
}

// Submit submit task
func Submit(player Player, packet *network.Message) {
	player.GetTaskData().GetTask(0).SetStatus(SUBMIT)
}
