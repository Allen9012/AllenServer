package cluster

import (
	"github.com/Allen9012/AllenGame/service"
	"net/rpc"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

var configDir = "./config/"

type SetupServiceFun func(s ...service.IService)

type NodeStatus int

const (
	Normal  NodeStatus = 0 //正常
	Discard NodeStatus = 1 //丢弃
)

type NodeInfo struct {
	NodeId            int
	NodeName          string
	Private           bool
	ListenAddr        string
	MaxRpcParamLen    uint32   //最大Rpc参数长度
	CompressBytesLen  int      //超过字节进行压缩的长度
	ServiceList       []string //所有的有序服务列表
	PublicServiceList []string //对外公开的服务列表
	DiscoveryService  []string //筛选发现的服务，如果不配置，不进行筛选
	NeighborService   []string
	status            NodeStatus
}

type NodeRpcInfo struct {
	nodeInfo NodeInfo
	client   *rpc.Client
}

var cluster Cluster

type Cluster struct {
	localNodeInfo           NodeInfo    //本结点配置信息
	masterDiscoveryNodeList []NodeInfo  //配置发现Master结点
	globalCfg               interface{} //全局配置

	localServiceCfg  map[string]interface{} //map[serviceName]配置数据*
	serviceDiscovery IServiceDiscovery      //服务发现接口

	locker         sync.RWMutex                //结点与服务关系保护锁
	mapRpc         map[int]NodeRpcInfo         //nodeId
	mapIdNode      map[int]NodeInfo            //NodeId:NodeInfo
	mapServiceNode map[string]map[int]struct{} //serviceName:map[NodeId]

	rpcServer                      rpc.Server
	rpcEventLocker                 sync.RWMutex        //Rpc事件监听保护锁
	mapServiceListenRpcEvent       map[string]struct{} //ServiceName
	mapServiceListenDiscoveryEvent map[string]struct{} //ServiceName
}

func (cls *Cluster) Init(localNodeId int, setupServiceFun SetupServiceFun) error {
	//1.初始化配置
	err := cls.InitCfg(localNodeId)
	if err != nil {
		return err
	}

	cls.rpcServer.Init(cls)
	cls.buildLocalRpc()

	//2.安装服务发现结点
	cls.SetupServiceDiscovery(localNodeId, setupServiceFun)
	service.RegRpcEventFun = cls.RegRpcEvent
	service.UnRegRpcEventFun = cls.UnRegRpcEvent
	service.RegDiscoveryServiceEventFun = cls.RegDiscoveryEvent
	service.UnRegDiscoveryServiceEventFun = cls.UnReDiscoveryEvent

	err = cls.serviceDiscovery.InitDiscovery(localNodeId, cls.serviceDiscoveryDelNode, cls.serviceDiscoverySetNodeInfo)
	if err != nil {
		return err
	}

	return nil
}

func GetCluster() *Cluster {
	return &cluster
}

func SetConfigDir(cfgDir string) {
	configDir = cfgDir
}

func (cls *Cluster) GetLocalNodeInfo() *NodeInfo {
	return &cls.localNodeInfo
}

func (cls *Cluster) Start() {
	cls.rpcServer.Start(cls.localNodeInfo.ListenAddr, cls.localNodeInfo.MaxRpcParamLen, cls.localNodeInfo.CompressBytesLen)
}
