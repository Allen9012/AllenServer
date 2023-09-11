package node

import (
	"errors"
	"fmt"
	"github.com/Allen9012/AllenGame/cluster"
	"github.com/Allen9012/AllenGame/console"
	"github.com/Allen9012/AllenGame/log"
	"github.com/Allen9012/AllenGame/service"
	"github.com/Allen9012/AllenGame/util/buildtime"
	"github.com/Allen9012/AllenGame/util/timer"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/9/8
	@desc:
	@modified by:
*/

var sig chan os.Signal
var nodeId int
var preSetupService []service.IService //预安装的服务
var profilerInterval time.Duration
var bValid bool
var configDir = "./config/"

type BuildOSType = int8

const (
	Windows BuildOSType = 0
	Linux   BuildOSType = 1
	Mac     BuildOSType = 2
)

func init() {
	sig = make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.Signal(10))

	console.RegisterCommandBool("help", false, "<-help> This help.", usage)
	console.RegisterCommandString("name", "", "<-name nodeName> Node's name.", setName)
	console.RegisterCommandString("start", "", "<-start nodeid=nodeid> Run originserver.", startNode)
	console.RegisterCommandString("stop", "", "<-stop nodeid=nodeid> Stop originserver process.", stopNode)
	console.RegisterCommandString("config", "", "<-config path> Configuration file path.", setConfigPath)
	//console.RegisterCommandString("console", "", "<-console true|false> Turn on or off screen log output.", openConsole)
	console.RegisterCommandString("loglevel", "debug", "<-loglevel debug|release|warning|error|fatal> Set loglevel.", setLevel)
	console.RegisterCommandString("logpath", "", "<-logpath path> Set log file path.", setLogPath)
	console.RegisterCommandInt("logsize", 0, "<-logsize size> Set log size(MB).", setLogSize)
	console.RegisterCommandInt("logchannelcap", 0, "<-logchannelcap num> Set log channel cap.", setLogChannelCapNum)
	console.RegisterCommandString("pprof", "", "<-pprof ip:port> Open performance analysis.", setPprof)
}

/* console config Handler */

func usage(val interface{}) error {
	ret := val.(bool)
	if ret == false {
		return nil
	}

	if len(buildtime.GetBuildDateTime()) > 0 {
		fmt.Fprintf(os.Stderr, "Welcome to Allen(build info: %s)\nUsage: AllenServer [-help] [-start node=1] [-stop] [-config path] [-pprof 0.0.0.0:6060]...\n", buildtime.GetBuildDateTime())
	} else {
		fmt.Fprintf(os.Stderr, "Welcome to Allen\nUsage: AllenServer [-help] [-start node=1] [-stop] [-config path] [-pprof 0.0.0.0:6060]...\n")
	}

	console.PrintDefaults()
	return nil
}

func setName(val interface{}) error {
	return nil
}

func stopNode(args interface{}) error {
	//1.解析参数
	param := args.(string)
	if param == "" {
		return nil
	}

	sParam := strings.Split(param, "=")
	if len(sParam) != 2 {
		return fmt.Errorf("invalid option %s", param)
	}
	if sParam[0] != "nodeid" {
		return fmt.Errorf("invalid option %s", param)
	}
	nodeId, err := strconv.Atoi(sParam[1])
	if err != nil {
		return fmt.Errorf("invalid option %s", param)
	}

	processId, err := getRunProcessPid(nodeId)
	if err != nil {
		return err
	}

	KillProcess(processId)
	return nil
}

func startNode(args interface{}) error {
	//1.解析参数
	param := args.(string)
	if param == "" {
		return nil
	}

	sParam := strings.Split(param, "=")
	if len(sParam) != 2 {
		return fmt.Errorf("invalid option %s", param)
	}
	if sParam[0] != "nodeid" {
		return fmt.Errorf("invalid option %s", param)
	}
	nodeId, err := strconv.Atoi(sParam[1])
	if err != nil {
		return fmt.Errorf("invalid option %s", param)
	}

	timer.StartTimer(10*time.Millisecond, 1000000)
	log.Info("Start running server.")
	//2.初始化node
	initNode(nodeId)

	//3.运行service
	service.Start()

	//4.运行集群
	cluster.GetCluster().Start()

	//5.记录进程id号
	writeProcessPid(nodeId)

	//6.监听程序退出信号&性能报告
	bRun := true
	var pProfilerTicker *time.Ticker = &time.Ticker{}
	if profilerInterval > 0 {
		pProfilerTicker = time.NewTicker(profilerInterval)
	}
	for bRun {
		select {
		case <-sig:
			log.Info("receipt stop signal.")
			bRun = false
		case <-pProfilerTicker.C:
			profiler.Report()
		}
	}
	cluster.GetCluster().Stop()
	//7.退出
	service.StopAllService()

	log.Info("Server is stop.")
	log.Close()
	return nil
}
func setConfigPath(val interface{}) error {
	configPath := val.(string)
	if configPath == "" {
		return nil
	}
	_, err := os.Stat(configPath)
	if err != nil {
		return fmt.Errorf("Cannot find file path %s", configPath)
	}

	cluster.SetConfigDir(configPath)
	configDir = configPath
	return nil
}

func setLevel(args interface{}) error {
	if args == "" {
		return nil
	}

	strlogLevel := strings.TrimSpace(args.(string))
	switch strlogLevel {
	case "trace":
		log.LogLevel = log.LevelTrace
	case "debug":
		log.LogLevel = log.LevelDebug
	case "info":
		log.LogLevel = log.LevelInfo
	case "warning":
		log.LogLevel = log.LevelWarning
	case "error":
		log.LogLevel = log.LevelError
	case "stack":
		log.LogLevel = log.LevelStack
	case "fatal":
		log.LogLevel = log.LevelFatal
	default:
		return errors.New("unknown level: " + strlogLevel)
	}
	return nil
}

func getRunProcessPid(nodeId int) (int, error) {
	f, err := os.OpenFile(fmt.Sprintf("%s_%d.pid", os.Args[0], nodeId), os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	pidByte, errs := io.ReadAll(f)
	if errs != nil {
		return 0, errs
	}

	return strconv.Atoi(string(pidByte))
}
func setLogPath(args interface{}) error {
	if args == "" {
		return nil
	}

	log.LogPath = strings.TrimSpace(args.(string))
	dir, err := os.Stat(log.LogPath) //这个文件夹不存在
	if err == nil && dir.IsDir() == false {
		return errors.New("Not found dir " + log.LogPath)
	}

	if err != nil {
		err = os.Mkdir(log.LogPath, os.ModePerm)
		if err != nil {
			return errors.New("Cannot create dir " + log.LogPath)
		}
	}

	return nil
}

func setLogSize(args interface{}) error {
	if args == "" {
		return nil
	}

	logSize, ok := args.(int)
	if ok == false {
		return errors.New("param logsize is error")
	}

	log.LogSize = int64(logSize) * 1024 * 1024

	return nil
}

func setLogChannelCapNum(args interface{}) error {
	if args == "" {
		return nil
	}

	logChannelCap, ok := args.(int)
	if ok == false {
		return errors.New("param logsize is error")
	}

	log.LogChannelCap = logChannelCap
	return nil
}

func setPprof(val interface{}) error {
	listenAddr := val.(string)
	if listenAddr == "" {
		return nil
	}

	go func() {
		err := http.ListenAndServe(listenAddr, nil)
		if err != nil {
			panic(fmt.Errorf("%+v", err))
		}
	}()

	return nil
}

/*	分割子函数 */

// initNode 初始化节点
func initNode(id int) {
	//1.初始化集群
	nodeId = id
	err := cluster.GetCluster().Init(GetNodeId(), Setup)
	if err != nil {
		log.Fatal("read system config is error ", log.ErrorAttr("error", err))
	}

	err = initLog()
	if err != nil {
		return
	}

	//2.顺序安装服务
	serviceOrder := cluster.GetCluster().GetLocalNodeInfo().ServiceList
	for _, serviceName := range serviceOrder {
		bSetup := false
		for _, s := range preSetupService {
			if s.GetName() != serviceName {
				continue
			}
			bSetup = true
			pServiceCfg := cluster.GetCluster().GetServiceCfg(s.GetName())
			s.Init(s, cluster.GetRpcClient, cluster.GetRpcServer, pServiceCfg)

			service.Setup(s)
		}

		if bSetup == false {
			log.Fatal("Service name " + serviceName + " configuration error")
		}
	}

	//3.service初始化
	service.Init()
}

// 初始化log
func initLog() error {
	if log.LogPath == "" {
		setLogPath("./log")
	}

	localnodeinfo := cluster.GetCluster().GetLocalNodeInfo()
	filepre := fmt.Sprintf("%s_%d_", localnodeinfo.NodeName, localnodeinfo.NodeId)
	logger, err := log.NewTextLogger(log.LogLevel, log.LogPath, filepre, true, log.LogChannelCap)
	if err != nil {
		fmt.Printf("cannot create log file!\n")
		return err
	}
	log.Export(logger)
	return nil
}

// GetNodeId 进程id号
func GetNodeId() int {
	return nodeId
}

// Setup setup service
func Setup(s ...service.IService) {
	for _, sv := range s {
		sv.OnSetup(sv)
		preSetupService = append(preSetupService, sv)
	}
}