package skill

import "github.com/Allen9012/AllenServer/business/module/scene/buff"

type Base struct {
	Id     uint32
	Desc   string
	Cd     int64
	Damage int64
	Buffs  []buff.Abstract
}
