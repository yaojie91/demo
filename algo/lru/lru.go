package lru

import (
	"container/list"
	"errors"
)

type CacheNode struct {
	K, V interface{}
}

// func (c *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
// 	return &CacheNode{k, v}
// }

type LRUCache struct {
	Cap      int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NEWLRUCache(cap int) *LRUCache {
	return &LRUCache{
		Cap:      cap,
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element),
	}
}

func (lru *LRUCache) Size() int {
	return lru.dlist.Len()
}

func (lru *LRUCache) Set(k, v interface{}) error {

	if lru.dlist == nil {
		return errors.New("LRU尚未初始化")
	}
	// 如果命中
	if pE, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pE)
		pE.Value.(*CacheNode).V = v
		return nil
	}
	// 如果没有命中
	newEliment := lru.dlist.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newEliment

	// 判断是否溢出
	if lru.dlist.Len() > lru.Cap {
		last := lru.dlist.Back()
		if last == nil {
			return nil
		}
		lastNode := last.Value.(*CacheNode)
		delete(lru.cacheMap, lastNode.K)
		lru.dlist.Remove(last)
	}
	return nil
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {

	if lru.cacheMap == nil {
		return v, false, errors.New("LRU尚未初始化")
	}

	if pEliment, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pEliment)
		return pEliment.Value.(*CacheNode).V, true, nil
	}
	return v, false, errors.New("没有找到")
}

func (lru *LRUCache) Remove(k interface{}) bool {

	if lru.cacheMap == nil {
		return false
	}

	if pE, ok := lru.cacheMap[k]; ok {
		cacheNode := pE.Value.(*CacheNode) // 断言CacheNode
		delete(lru.cacheMap, cacheNode.K)  // 删除map键值对
		lru.dlist.Remove(pE)
		return true
	}
	return false
}
