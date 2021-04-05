package main

import "container/list"

type LRUCache struct {
	dataMap  map[int]*list.Element
	dataList list.List
	cap      int
}

type cacheData struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dataMap:  make(map[int]*list.Element, capacity),
		cap:      capacity,
		dataList: list.List{},
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.dataMap[key]
	if !ok {
		return -1
	}
	this.dataList.MoveToFront(value)
	return value.Value.(cacheData).value
}

func (this *LRUCache) Put(key int, value int) {
	if oldValue, ok := this.dataMap[key]; ok {
		tmpData := cacheData{
			value: value,
			key:   key,
		}
		oldValue.Value = tmpData
		this.dataList.MoveToFront(oldValue)
		return
	}

	if this.dataList.Len() >= this.cap {
		backData := this.dataList.Back()
		this.dataList.Remove(backData)
		delete(this.dataMap, backData.Value.(cacheData).key)
	}

	ele := this.dataList.PushFront(cacheData{value: value, key: key})
	this.dataMap[key] = ele
}
