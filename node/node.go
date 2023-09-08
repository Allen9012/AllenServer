package node

import (
	"os"
	"os/signal"
	"syscall"
)

/*
*

	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/9/8
	@desc:
	@modified by:

*
*/
var sig chan os.Signal

func init() {
	sig = make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.Signal(10))

	console.RegisterCommandBool("help", false, "<-help> This help.", usage)
	console.RegisterCommandString("name", "", "<-name nodeName> Node's name.", setName)
	console.RegisterCommandString("start", "", "<-start nodeid=nodeid> Run originserver.", startNode)
	console.RegisterCommandString("stop", "", "<-stop nodeid=nodeid> Stop originserver process.", stopNode)
	console.RegisterCommandString("config", "", "<-config path> Configuration file path.", setConfigPath)
	console.RegisterCommandString("console", "", "<-console true|false> Turn on or off screen log output.", openConsole)
	console.RegisterCommandString("loglevel", "debug", "<-loglevel debug|release|warning|error|fatal> Set loglevel.", setLevel)
	console.RegisterCommandString("logpath", "", "<-logpath path> Set log file path.", setLogPath)
	console.RegisterCommandInt("logsize", 0, "<-logsize size> Set log size(MB).", setLogSize)
	console.RegisterCommandInt("logchannelcap", 0, "<-logchannelcap num> Set log channel cap.", setLogChannelCapNum)
	console.RegisterCommandString("pprof", "", "<-pprof ip:port> Open performance analysis.", setPprof)
}
