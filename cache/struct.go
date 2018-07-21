package cache

import (
	"container/list"
	"sync"
)

type Cache struct {
	//limit size
	limit int
	//items will expire after defaultExpireInterval
	defaultExpireInterval int64
	maxExpireTime         int64
	items                 map[string]*list.Element
	//oldest list
	list *list.List
	lock *sync.RWMutex
	onRemove func(key string ,val interface{})
}

type item struct {
	//expire time in unix timestamp
	ExpireTime int64
	object interface{}
}