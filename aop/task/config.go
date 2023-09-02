package task

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/1
  @desc:
  @modified by:
**/

type Config struct {
	Id       uint32        `json:"id"`
	Name     string        `json:"name"`
	DropId   uint32        `json:"dropId"` //
	Category int           `json:"category"`
	Targets  []*TargetConf `json:"targets"`
}

type TargetConf struct {
}
