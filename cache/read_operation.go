package cache

import (
	"time"
	"container/list"
)

/*
	Get object from cache
	1 lock
	2 get
	3 check is expired
	4 update elimination list
	5 return
	6 unlock
*/
func (c *cache)Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	elmt := c.getAndCheckExpire(key)
	if elmt == nil {
		return nil
	}
	c.updateList(elmt)
	return elmt.Value
}

func (c *cache)updateList(elmt *list.Element)  {
	c.list.MoveToFront(elmt)
}

//unsafe basic operation
func (c *cache)getAndCheckExpire(key string) *list.Element  {
	itm , ok := c.items[key]
	if !ok {
		return nil
	}
	//expired
	if time.Now().Unix() > itm.Value.(item).ExpireTime {
		//remove item
		delete(c.items, key)
		c.list.Remove(itm)
		return nil
	}
	return itm
}