package example

import (
	"fmt"
	"github.com/Allen9012/AllenServer/business/module/condition"
	"math"
	"testing"
)

func TestEvent(t *testing.T) {
	te := TEvent{
		Subscribers: make([]condition.Condition, 0),
	}
	tg := &TTarget{
		Id:   111,
		Data: 1,
	}
	te.Attach(tg)
	te.Data = 1
	te.Notify()
	fmt.Println("CheckArrived:", tg.CheckArrived())
}

func TestTask(t *testing.T) {
	NewTTask(nil)
	dp := make([]int, 10, math.MinInt)
	print(dp)
}
