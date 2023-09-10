package network

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc: Agent
  @modified by:
**/

type Agent interface {
	Run()
	OnClose()
}
