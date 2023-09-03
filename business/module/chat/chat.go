package chat

import (
	"container/ring"
	"go.mongodb.org/mongo-driver/bson"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type System struct {
	latestOnlineMessages    *ring.Ring
	latestCrossZoneMessages *ring.Ring
	latestZoneMessages      *ring.Ring
	latestCrossSrvMessages  *ring.Ring
	Owner
}

type Chat struct {
	Id      uint64
	Content string
}

func (c *Chat) ToDB() *Model {
	return &Model{
		Id:      c.Id,
		Content: c.Content,
	}
}

func (c *System) SetOwner(owner Owner) {
	c.Owner = owner
}

func (c *System) ToDB() bson.M {

	return nil
}
