package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Copyright © 2023 github.com/Allen9012 All rights reserved.
@author: Allen
@since: 2023/8/27
@desc:  模拟客户端收发请求
@modified by:
*/

type ClientConsole struct {
	chInput chan *InputParam
}

type InputParam struct {
	Command string
	Param   []string
}

func NewClientConsole() *ClientConsole {
	c := &ClientConsole{}
	return c
}

func (c *ClientConsole) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input err ,check your input and  try again !!!")
			continue
		}
		split := strings.Split(readString, " ")
		if len(split) == 0 {
			fmt.Println("input err, check your input and  try again !!! ")
			continue
		}
		in := &InputParam{
			Command: split[0],
			Param:   split[1:],
		}
		c.chInput <- in
	}
}
