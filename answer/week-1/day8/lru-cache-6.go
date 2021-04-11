package day8

import (
	"container/list"
)

type LRUCache struct {
	max   int
	ll    list.List
	catch map[int]*list.Element
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
	if v, ok := this.catch[key]; ok {
		this.ll.MoveToFront(v)
		return v.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.catch == nil {
		this.catch = make(map[int]*list.Element)
		this.ll = list.List{}
	}
	if v, ok := this.catch[key]; ok {
		this.ll.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}
	element := this.ll.PushFront(&entry{key, value})
	this.catch[key] = element
	if this.ll.Len() > this.max {
		rmv := this.ll.Remove(this.ll.Back())
		delete(this.catch, rmv.(*entry).key)
	}
}
