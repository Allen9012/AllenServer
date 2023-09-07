package friendevent

import (
	"github.com/Allen9012/AllenServer/internal/event"
)

type AddOrDelFriendEvent struct {
	CurFriendCount int
	event.Base
}

func (e *AddOrDelFriendEvent) GetDesc() string {
	return ""
}
