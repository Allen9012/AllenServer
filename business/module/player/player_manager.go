package player

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

func NewPlayerMgr() *Manager {
	return &Manager{
		players: make(map[uint64]*Player),
		addPCh:  make(chan *Player, 1),
	}
}

// Add ...
func (pm *Manager) Add(p *Player) {
	if pm.players[p.UID] != nil {
		return
	}
	pm.players[p.UID] = p
	go p.Start()
}

// Del ...
func (pm *Manager) Del(p Player) {
	delete(pm.players, p.UID)
}

func (pm *Manager) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		}
	}
}

func (pm *Manager) GetPlayer(uId uint64) *Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
