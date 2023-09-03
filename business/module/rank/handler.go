package rank

import (
	"github.com/Allen9012/AllenServer/business/module/player"
	"github.com/Allen9012/AllenServer/network"
)

func init() {
	register.Register(222, GetRankList)

}

func GetRankList(player *player.Player, packet *network.Packet) {

}
