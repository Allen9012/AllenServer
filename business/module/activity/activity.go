package activity

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: 抽象活动方法
  @modified by:
**/

type Activity interface {
	CheckInTimeRange() bool
}
