package cache

import (
	"container/list"
	"sync"
)

func (c *Cache)updateList(elmt *list.Element)  {
	c.list.MoveToFront(elmt)
}

func NewCache(limit int , defaultExpireInterval int64, onRemove func(key string , value interface{})) (*Cache , error) {
	var maxExpireTime int64
	maxExpireTime = 1<<63 - 1 // biggest int
	if defaultExpireInterval < 0 {
		defaultExpireInterval = maxExpireTime
	}
	return &Cache{
		limit,
		defaultExpireInterval ,
		maxExpireTime,make(map[string]*list.Element),
		list.New(),
		new(sync.RWMutex),
		onRemove,
	},nil
}