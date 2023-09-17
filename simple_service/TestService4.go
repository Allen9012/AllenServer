package simple_service

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: 监听节点链接断开
  @modified by:
**/
import (
	"fmt"
	"github.com/Allen9012/AllenGame/node"
	"github.com/Allen9012/AllenGame/service"
)

func init() {
	node.Setup(&TestService4{})
}

type TestService4 struct {
	service.Service
}

func (slf *TestService4) OnInit() error {
	fmt.Println("【监听节点链接断开服务】启动")
	return nil
}

// 监听节点链接
func (slf *TestService4) OnNodeConnected(nodeId int) {
	fmt.Printf("nodeId 链接: %v\n", nodeId)
}

// 监听节点断开
func (slf *TestService4) OnNodeDisconnect(nodeId int) {
	fmt.Printf("nodeId 断开: %v\n", nodeId)
}
