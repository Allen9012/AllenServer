package actor

type Monster struct {
	*Base
	real MonsterReal
}

func NewMonster() *Monster {
	return &Monster{
		Base: &Base{
			Hp:     0,
			Damage: 0,
		}}
}

func (m *Monster) OnDamage(delta int64) {
	m.Hp -= delta
}

func (m *Monster) Attack() {
	//TODO implement me
	panic("implement me")
}

func (m *Monster) OnMove() {
	//TODO implement me
	panic("implement me")
}
