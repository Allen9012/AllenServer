package launch

import (
	"github.com/Allen9012/AllenServer/aop/logger"
	"github.com/Allen9012/AllenServer/business/server/world"
	"github.com/Allen9012/sugar"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

func main() {
	world.MM = world.NewMgrMgr()
	go world.MM.Start()
	logger.Info("server start !!")
	sugar.WaitSignal(world.MM.OnSystemSignal)
}
