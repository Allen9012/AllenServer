package simple_module

import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
)

func init() {
	node.Setup(&TestServiceModel{})
}

type TestServiceModel struct {
	service.Service
}

func (slf *TestServiceModel) OnInit() error {
	fmt.Printf("【插入model服务】启动\n")

	//两种插入module方式
	//1,平行插入
	//2,深度插入

	//新建两个Module对象
	module1 := &Module1{}
	module2 := &Module2{}
	module3 := &Module3{}

	//将module1添加到服务中
	module1Id, _ := slf.AddModule(module1)
	//将module3添加到服务中
	module3Id, _ := slf.AddModule(module3)

	//在module1中添加module2模块
	module1.AddModule(module2)

	fmt.Printf("module1Id: %v\n", module1Id)
	fmt.Printf("module3Id: %v\n", module3Id)

	fmt.Printf("module2Id: %v\n", module2.GetModuleId())

	//释放模块module1
	slf.ReleaseModule(module1Id)
	//释放模块module3
	slf.ReleaseModule(module3Id)

	return nil
}
