package bucket

import (
	"math"
)

type Item struct {
	Key   int
	Value string
}
type MyBucket struct {
	items []Item
}

func NewBucket() MyBucket {
	b := MyBucket{items: make([]Item, 0, 10e6)}
	for i := 0; i < 1e6; i++ {
		b.items = append(b.items, Item{
			Key:   i,
			Value: "value",
		})
	}
	return b
}

func (p *MyBucket) Size() int {
	return len(p.items)
}
func (p *MyBucket) Get(key int) (Item, bool) {
	l := len(p.items)
	low := float64(0)
	high := float64(l - 1)
	for low <= high {
		mid := math.Floor((high + low) / 2)
		item := p.items[int(mid)]
		if item.Key == key {
			return item, true
		}
		if key < item.Key {
			high = mid - 1
		}
		if key > item.Key {
			low = mid + 1
		}
	}
	return Item{}, false
}
