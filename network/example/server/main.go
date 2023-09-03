package main

import "github.com/Allen9012/AllenServer/network"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func main() {
	server := network.NewServer(":8888", "tcp")
	server.Run()
	select {}
}
