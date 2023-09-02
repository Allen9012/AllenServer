package algorithm

import "math/rand"

/*
	常规实现的话是开辟一个slice，可能还需要删除随机出来的元素，还需要再开辟slice随机，这就非常影响性能
	下面的实现方法就比较巧妙，不需要删除元素，也不需要再开辟slice，洗牌算法
*/

// Samples returns N random unique items from collection.
func Samples[T any](collection []T, count int) []T {
	size := len(collection)

	ts := append([]T{}, collection...)

	results := []T{}

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i

		index := rand.Intn(size - i)
		results = append(results, ts[index])

		// Removes element.
		// It is faster to swap with last element and remove it.
		ts[index] = ts[copyLength-1]
		ts = ts[:copyLength-1]
	}

	return results
}
