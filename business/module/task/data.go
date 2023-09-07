package task

import "sync"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type Data struct {
	Tasks        sync.Map
	Achievements sync.Map
	taskCache    sync.Map // map[EventCategory][]uint64
}

func NewTaskData() *Data {
	return &Data{
		Tasks:        sync.Map{},
		Achievements: sync.Map{},
		taskCache:    sync.Map{},
	}
}

func (d *Data) ToDB() {

}

func (d *Data) LoadFromDB() {

}

func (d *Data) GetTask(id uint64) Task {
	value, ok := d.Tasks.Load(id)
	if ok {
		return value.(Task)
	}
	return nil
}

func (d *Data) SyncAllTasks(player Player) {
	//todo
}

func (d *Data) AddTaskCache(category EventCategory, taskId uint64) {
	value, ok := d.taskCache.Load(category)
	uint64s := make([]uint64, 0)
	if ok { // 已经有
		uint64s = value.([]uint64)
	}
	uint64s = append(uint64s, taskId)
	d.taskCache.Store(category, uint64s)
}

func (d *Data) GetTaskCache(category EventCategory) []uint64 {
	if value, ok := d.taskCache.Load(category); ok {
		return value.([]uint64)
	}
	return nil
}
