package function

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/27
  @desc:
  @modified by:
**/

func CheckInNumberSlice[T uint64](id T, slice []T) bool {
	for _, v := range slice {
		if id == v {
			return true
		}
	}
	return false
}

// DelEleInSlice 删除给定元素
func DelEleInSlice[T uint64](id T, slice []T) []T {
	for _, v := range slice {
		if id == v {

		}
	}
}
