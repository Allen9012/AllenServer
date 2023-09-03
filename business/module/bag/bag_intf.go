package bag

import "github.com/Allen9012/AllenServer/business/module/bag/item"

type Bag interface {
	AddItem(item item.Item)
	DelItem(item item.Item)
}
