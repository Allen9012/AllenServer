package scene

import (
	"github.com/Allen9012/AllenServer/business/module/scene/actor"
	"google.golang.org/protobuf/proto"
)

type Abstract interface {
	OnCreate()
	Run()
	OnDestroy()
	loop()
	monitor()
}

type Notify interface {
	NotifyAll(message proto.Message)
	NotifyNearby(actor actor.Actor, message proto.Message)
	NotifyPlayer(playerId uint64, message proto.Message)
}

type Action interface {
	OnNextWave()
	OnMonsterDie()
	OnWaveEnd()
}

type FightScene interface {
	Abstract
	Action
}
