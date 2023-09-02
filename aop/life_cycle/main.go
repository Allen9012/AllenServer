package life_cycle

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc: life cycle of game server
  @modified by:
**/

func main() {
	LoadConf()
	LoadDBData()
	StartGameServer()
	ResolveGameLogic()
	//	停服更新
	StopGameServer()
}

// StartGameServer 启动游戏服
func StartGameServer() {
	Log()
}

// LoadConf 加载配置
func LoadConf() {
	Log()
}

// LoadDBData 加载数据库
func LoadDBData() {
	Log()
}

// ResolveGameLogic 处理游戏逻辑 - 客户端发来的请求
func ResolveGameLogic() {
	SaveGameData()
	Log()
}

// SaveGameData 存储游戏数据
func SaveGameData() {
	Log()
}

// StopGameServer 停止游戏服
func StopGameServer() {
	Log()
}

// Log 打印日志
func Log() {

}
