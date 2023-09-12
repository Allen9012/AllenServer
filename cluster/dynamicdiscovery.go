package cluster

import (
	"github.com/Allen9012/AllenGame/rpc"
	"github.com/Allen9012/AllenGame/service"
)

const DynamicDiscoveryMasterName = "DiscoveryMaster"
const DynamicDiscoveryClientName = "DiscoveryClient"
const RegServiceDiscover = DynamicDiscoveryMasterName + ".RPC_RegServiceDiscover"
const SubServiceDiscover = DynamicDiscoveryClientName + ".RPC_SubServiceDiscover"
const AddSubServiceDiscover = DynamicDiscoveryMasterName + ".RPC_AddSubServiceDiscover"

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

func init() {
	masterService.SetName(DynamicDiscoveryMasterName)
	clientService.SetName(DynamicDiscoveryClientName)
}

func (d *DynamicDiscoveryClient) InitDiscovery(localNodeId int, funDelNode FunDelNode, funSetNodeInfo FunSetNodeInfo) error {
	//TODO implement me
	panic("implement me")
}

func (d *DynamicDiscoveryClient) OnNodeStop() {
	//TODO implement me
	panic("implement me")
}

func getDynamicDiscovery() IServiceDiscovery {
	return &clientService
}
