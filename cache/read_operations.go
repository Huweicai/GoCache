package cache

import (
	"container/list"
	"time"
)

/*
	Get object from Cache
	1 lock
	2 get
	3 check is expired
	4 update elimination list
	5 return
	6 unlock
*/
func (c *Cache) Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	elmt := c.getAndCheckExpire(key)
	if elmt == nil {
		return nil
	}
	c.updateList(elmt)
	return elmt.Value.(*item).object
}

/**
unsafe basic operation
*/
func (c *Cache) getAndCheckExpire(key string) *list.Element {
	itm, ok := c.items[key]
	if !ok {
		return nil
	}
	//expired
	if time.Now().Unix() > itm.Value.(*item).ExpireTime {
		//remove item
		delete(c.items, key)
		c.list.Remove(itm)
		c.onRemove(key, itm.Value)
		return nil
	}
	return itm
}

/*
Get returns default value while get nothing
*/
func (c *Cache) DefaultGet(key string, defaultVal interface{}) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	elmt := c.getAndCheckExpire(key)
	if elmt == nil {
		return defaultVal
	}
	c.updateList(elmt)
	return elmt.Value.(*item).object
}

