package concurrent

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc:
  @modified by:
**/

type task struct {
	queueId int64
	fn      func() bool
	cb      func(err error)
}

type worker struct {
	*dispatch
}
