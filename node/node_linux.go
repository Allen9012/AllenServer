//go:build linux
// +build linux

package node

import (
	"fmt"
	"syscall"
)

/*
	Copyright © 2023 github.com/Allen9012 All rights reserved.
	@author: Allen
	@since: 2023/9/8
	@desc:
	@modified by:
*/

func KillProcess(processId int) {
	err := syscall.Kill(processId, syscall.Signal(10))
	if err != nil {
		fmt.Printf("kill processid %d is fail:%+v.\n", processId, err)
	} else {
		fmt.Printf("kill processid %d is successful.\n", processId)
	}
}

func GetBuildOSType() BuildOSType {
	return Linux
}
