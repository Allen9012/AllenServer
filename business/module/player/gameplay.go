package player

import (
	"github.com/Allen9012/AllenServer/business/module/bag"
	"github.com/Allen9012/AllenServer/business/module/building"
	"github.com/Allen9012/AllenServer/business/module/chat"
	"github.com/Allen9012/AllenServer/business/module/email"
	"github.com/Allen9012/AllenServer/business/module/friend"
	"github.com/Allen9012/AllenServer/business/module/pet"
	"github.com/Allen9012/AllenServer/business/module/plant"
	"github.com/Allen9012/AllenServer/business/module/shop"
	"github.com/Allen9012/AllenServer/business/module/task"
	"github.com/Allen9012/AllenServer/business/module/vip"
)

type GamePlay struct {
	friendSystem   *friend.System
	privateChat    *chat.PrivateChat
	taskData       *task.Data
	petSystem      *pet.System
	shopData       *shop.Data
	bagSystem      *bag.System
	vip            *vip.Vip
	buildingSystem *building.System
	plantSystem    *plant.System
	emailData      *email.Data
}

func InitGamePlay() GamePlay {
	return GamePlay{
		friendSystem:   nil,
		privateChat:    nil,
		taskData:       nil,
		petSystem:      nil,
		shopData:       nil,
		bagSystem:      nil,
		vip:            nil,
		buildingSystem: nil,
	}
}

func (p *GamePlay) GetTaskData() *task.Data {
	return p.taskData
}

func (p *GamePlay) GetPetSystem() *pet.System {
	return p.petSystem
}

func (p *GamePlay) GetShopData() *shop.Data {
	return p.shopData
}

func (p *GamePlay) GetBagSystem() *bag.System {
	return p.bagSystem
}

func (p *GamePlay) GetVip() *vip.Vip {
	return p.vip
}

func (p *GamePlay) GetBuildingSystem() *building.System {
	return p.buildingSystem
}

func (p *GamePlay) GetPlantSystem() *plant.System {
	return p.plantSystem
}

func (p *GamePlay) GetEmailData() *email.Data {
	return p.emailData
}
