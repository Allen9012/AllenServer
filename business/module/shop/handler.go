package shop

import (
	"errors"
	"github.com/Allen9012/AllenServer/network"
	"github.com/Allen9012/AllenServer/protos/gen/messageID"
	"sync"
)

type Handler struct {
	Id messageID.MessageId
	Fn func(player IPlayer, packet *network.Message)
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
		handlers[0] = &Handler{
			0,
			BuyItem,
		}
	})
}

// BuyItem  buy  item
func BuyItem(player IPlayer, packet *network.Message) {
	player.GetShopData()
}
