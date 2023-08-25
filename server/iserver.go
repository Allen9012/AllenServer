package server

// IService is the base interface for all services.
type IService interface {
	Start()
	Reload()
	Init(config interface{}, processId int)
	Stop()
}
