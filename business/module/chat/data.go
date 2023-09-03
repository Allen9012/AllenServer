package chat

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type Model struct {
	Id      uint64 `bson:"id"`
	Content string `bson:"content"`
}
