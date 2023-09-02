package event

import "github.com/Allen9012/AllenServer/business/module/condition"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type Event interface {
	Notify()
	Attach(condition condition.Condition)
	Detach(id uint32)
}
