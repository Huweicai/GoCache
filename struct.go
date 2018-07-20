package GoCache

import (
	"time"
	"container/list"
	"sync"
)

type Cache struct {
	//limit size
	limit int
	items map[string]list.Element
	//oldest list
	list list.List
	lock  sync.RWMutex
}

type Item struct {
	//expire time in unix nano timestamp
	ExpireTime time.Duration
	Object struct{}
}