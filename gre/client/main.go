package main

import "github.com/Allen9012/sugar"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func main() {
	c := NewClient()
	c.InputHandlerRegister()
	c.Run()
	sugar.WaitSignal(c.OnSystemSignal)
}
