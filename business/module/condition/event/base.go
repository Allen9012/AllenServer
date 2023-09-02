package event

import "github.com/Allen9012/AllenServer/business/module/condition"

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/2
  @desc:
  @modified by:
**/

type Base struct {
	Subscribers []condition.Condition
}

func (b *Base) Notify() {
	for _, subscriber := range b.Subscribers {
		subscriber.OnNotify(b)
	}
}

func (b *Base) Attach(c condition.Condition) {
	b.Subscribers = append(b.Subscribers, c)
}

func (b *Base) Detach(id uint32) {
	for i, subscriber := range b.Subscribers {
		if subscriber.GetId() == id {
			b.Subscribers = append(b.Subscribers[:i], b.Subscribers[i+1:]...)
		}
	}
}
