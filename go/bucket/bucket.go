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
func (p *MyBucket) Get(key int) (Item, int, bool) {
	l := len(p.items)
	low := 0
	high := l - 1
	i := 0
	for low <= high {
		i++
		mid := int(math.Floor(float64(high+low) / 2))
		item := p.items[mid]
		if item.Key == key {
			return item, i, true
		}
		if key < item.Key {
			high = mid - 1
		}
		if key > item.Key {
			low = mid + 1
		}
	}
	return Item{}, i, false
}
