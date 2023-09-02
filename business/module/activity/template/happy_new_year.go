package template

import "time"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type HappyNewYear struct {
	ID        uint32
	StartTime time.Time
	EndTime   time.Time
}

func (h *HappyNewYear) Init(conf Conf) *HappyNewYear {
	return &HappyNewYear{}
}

func (h *HappyNewYear) GetReward() {

}
