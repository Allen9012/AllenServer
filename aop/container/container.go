package container

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc: 容器思想管理
  @modified by:
**/

type IDelegate interface {
	Save(query, update interface{})
	Set(tag string, val interface{})
	Get(tag string) interface{}
}

type IContainer interface {
	IDelegate
	Add(interface{})
	Del(interface{})
	GetItem(val interface{}) interface{}
	SetItem(val interface{}, items interface{})
}
