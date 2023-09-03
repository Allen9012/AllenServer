package world

import (
	"github.com/Allen9012/AllenServer/business/module/activity"
	"github.com/Allen9012/AllenServer/business/module/bag"
	"github.com/Allen9012/AllenServer/business/module/chat"
	"github.com/Allen9012/AllenServer/business/module/email"
	"github.com/Allen9012/AllenServer/business/module/friend"
	"github.com/Allen9012/AllenServer/business/module/minigame"
	"github.com/Allen9012/AllenServer/business/module/rank"
	"github.com/Allen9012/AllenServer/business/module/task"
)

type GamePlay struct {
	activity activity.Abstract
	bag      bag.Abstract
	chat     chat.Abstract
	rank     rank.Abstract
	email    email.Abstract
	friend   friend.Abstract
	minigame minigame.Abstract
	task     task.Abstract
}
type Option func(play *GamePlay) *GamePlay

func WithActivity(activity activity.Abstract) Option {
	return func(play *GamePlay) *GamePlay {
		play.activity = activity
		return play
	}
}

func NewGamePlay(option ...Option) *GamePlay {
	g := &GamePlay{}
	for _, op := range option {
		op(g)
	}
	return g
}
