package mongo

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type Player struct {
	Uid        uint64   `bson:"uid"`
	NickName   string   `bson:"nickName"`
	Sex        int      `bson:"sex"`
	FriendList []uint64 `bson:"friendList"`
}
