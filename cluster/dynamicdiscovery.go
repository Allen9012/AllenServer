package cluster

import (
	"github.com/Allen9012/AllenGame/rpc"
	"github.com/Allen9012/AllenGame/service"
)

type DynamicDiscoveryMaster struct {
	service.Service

	mapNodeInfo map[int32]struct{}
	nodeInfo    []*rpc.NodeInfo
}

type DynamicDiscoveryClient struct {
	service.Service

	funDelService FunDelNode
	funSetService FunSetNodeInfo
	localNodeId   int

	mapDiscovery map[int32]map[int32]struct{} //map[masterNodeId]map[nodeId]struct{}
}

var masterService DynamicDiscoveryMaster
var clientService DynamicDiscoveryClient
