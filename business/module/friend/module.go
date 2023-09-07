package friend

import (
	"github.com/Allen9012/AllenServer/aop/module_router"
	"github.com/Allen9012/AllenServer/internal"
	"github.com/Allen9012/AllenServer/protos/gen/module"
	"sync"
)

var (
	Mod         *Module
	onceInitMod sync.Once
)

func init() {
	internal.ModuleManager.RegisterModule(module.Module_Friend.String(), GetMod())
}

type Module struct {
	*internal.BaseModule
}

func GetMod() *Module {
	Mod = &Module{internal.NewBaseModule()}

	return Mod
}

func (m *Module) GetName() string {
	return module.Module_Friend.String()
}

func (m *Module) RegisterHandler() {
	module_router.RegisterModuleMessageHandler(module.Module_Friend, 0, nil)
}
