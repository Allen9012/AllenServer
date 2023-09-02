package transport

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/30
  @desc: 存一些基础数据
  @modified by:
**/

type Player struct {
	UID        uint64   `bson:"uid"`
	NickName   string   `bson:"nickName"`
	Sex        int      `bson:"sex"`
	FriendList []uint64 `bson:"friendList"`
}
