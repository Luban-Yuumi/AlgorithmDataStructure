package day8

import (
	"container/list"
)

type LRUCache struct {
	max   int
	ll    list.List
	cache map[int]*list.Element
}
type entry struct {
	key   int
	value int
}

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		capacity,
		list.List{},
		make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.ll.MoveToFront(v)
		return v.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache == nil {
		this.cache = make(map[int]*list.Element)
		this.ll = list.List{}
	}
	if v, ok := this.cache[key]; ok {
		this.ll.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}
	element := this.ll.PushFront(&entry{key, value})
	this.cache[key] = element
	if this.ll.Len() > this.max {
		rmv := this.ll.Remove(this.ll.Back())
		delete(this.cache, rmv.(*entry).key)
	}
}
