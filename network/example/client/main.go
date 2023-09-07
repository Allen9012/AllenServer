package main

import "github.com/Allen9012/AllenServer/network"

/*
*

	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/8/27
	@desc:
	@modified by:

*
*/
func main() {
	// 表示请求这个地址
	client := network.NewClient(":8888")
	client.Run()
	select {}
}
