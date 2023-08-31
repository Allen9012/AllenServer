package demo

import (
	"github.com/Allen9012/AllenServer/demo/logger"
	"github.com/Allen9012/AllenServer/demo/world"
	"github.com/Allen9012/sugar"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/30
	@desc:
	@modified by:
*/

func main() {
	world.MM = world.NewMgrMgr()
	go world.MM.Run()
	select {}
	logger.Info("server start !!")
	sugar.WaitSignal(world.MM.OnSystemSignal)
}
