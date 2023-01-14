package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	li, exist := lc.items[key]
	if exist {
		lc.items[key].Value.(*cacheItem).value = value

		lc.queue.MoveToFront(li)
	} else {
		li := lc.queue.PushFront(&cacheItem{
			key:   key,
			value: value,
		})
		lc.items[key] = li
		if lc.queue.Len() > lc.capacity {
			delete(lc.items, lc.queue.Back().Value.(*cacheItem).key)

			lc.queue.Remove(lc.queue.Back())
		}
	}

	return exist
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	li, exist := lc.items[key]
	if exist {
		lc.queue.MoveToFront(li)
		lc.items[key] = lc.queue.Front()

		return lc.items[key].Value.(*cacheItem).value, exist
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.queue = NewList()
	lc.items = make(map[Key]*ListItem, lc.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
