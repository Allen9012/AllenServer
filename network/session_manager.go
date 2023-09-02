package network

import "sync"

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc: 维护session
	@modified by:
*/

type SessionMgr struct {
	Sessions map[uint64]*Session
	Counter  int64 //计数器
	Mutex    sync.Mutex
	Pid      int64
}

var (
	SessionMgrInstance SessionMgr
	onceInitSessionMgr sync.Once
)

// 单例初始化
func init() {
	onceInitSessionMgr.Do(func() {
		SessionMgrInstance = SessionMgr{
			Sessions: make(map[uint64]*Session),
			Counter:  0,
			Mutex:    sync.Mutex{},
		}
	})
}

// AddSession ...
func (sm *SessionMgr) AddSession(s *Session) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()
	if val := sm.Sessions[s.UID]; val != nil {
		if val.IsClose {
			sm.Sessions[s.UID] = s
		} else {
			return
		}
	}
}

// DelSession ...
func (sm *SessionMgr) DelSession(UID uint64) {
	delete(sm.Sessions, UID)
}
