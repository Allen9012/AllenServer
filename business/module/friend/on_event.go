package friend

import (
	"github.com/Allen9012/AllenServer/internal"
	"github.com/Allen9012/AllenServer/internal/event"
)

type EventHandle func(iEvent event.IEvent)

type EventWrap struct {
	IPlayer
	event.IEvent
}

func (m *Module) OnEvent(c internal.Character, event event.IEvent) {
	//TODO implement me
	panic("implement me")
}

func (m *Module) SetEventCategoryActive(eventCategory int) {
	//TODO implement me
	panic("implement me")
}
