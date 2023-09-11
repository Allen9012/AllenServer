package cluster

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:	静态配置服务发现，启动之后默认配置，不会修改
  @modified by:
**/

type ConfigDiscovery struct {
	funDelService FunDelNode
	funSetService FunSetNodeInfo
	localNodeId   int
}

func (discovery *ConfigDiscovery) InitDiscovery(localNodeId int, funDelNode FunDelNode, funSetNodeInfo FunSetNodeInfo) error {
	discovery.localNodeId = localNodeId
	discovery.funDelService = funDelNode
	discovery.funSetService = funSetNodeInfo

	//解析本地其他服务配置
	_, nodeInfoList, err := GetCluster().readLocalClusterConfig(0)
	if err != nil {
		return err
	}

	for _, nodeInfo := range nodeInfoList {
		if nodeInfo.NodeId == localNodeId {
			continue
		}

		discovery.funSetService(&nodeInfo)
	}

	return nil
}

func (discovery *ConfigDiscovery) OnNodeStop() {
}
