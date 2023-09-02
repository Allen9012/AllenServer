package template

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type Conf struct {
	ID          uint32
	Description string
	StartTime   string
	EndTime     string
	Reward      string
	Category    string
	Param1      string
	Param2      string
	Param3      string
}

// Verify 表格检查
func (c *Conf) Verify() {
	//	一致性检查
	//	关联性检查
	//	字段是否必须存在
	//	类型检查
	//	ASR检查
}

func (c *Conf) AfterAllVerify() {
	//	业务逻辑相关检查，可以在所有配置都加载完之后，执行这个方法
}
