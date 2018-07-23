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

/*
Remove all the items
 */
func (c *Cache)Truncate()  {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items = make(map[string]*list.Element)
}

func (c *Cache)RemoveOldest()(key string , val interface{})  {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.removeOldest()
	return
}


//unsafe inner removeOldest
func (c *Cache)removeOldest()  {
	end := c.list.Back()
	c.list.Remove(end)
	key := end.Value.(*item).key
	val := end.Value.(*item).object
	delete(c.items , key)
	c.onRemove(key , val)
}

func (c *Cache)size()(s int) {
	s = len(c.items)
	return
}

func (c *Cache)Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()

}