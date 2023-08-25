package main

import (
	"AllenServer/server"
	"AllenServer/utils/logger"
)

func main() {
	server.Oasis = server.NewWorld()
	go server.Oasis.Start()
	logger.Info("server start !!")
	sugar.WaitSignal(server.Oasis.OnSystemSignal)
}
