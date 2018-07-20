package cache

/*
	Set add or update an item to cache
	1 check expireTime
	2 check if
	item expires after expireTime when it's set positive,
	but if expireTime equals zero , it uses default expireTime,
	and item never expires while it's negative.
*/
func (c *cache)Set(key string , val interface{} , expireTime int64)  {
	c.lock.Lock()
	defer c.lock.Unlock()
	if expireTime == 0 {
		expireTime = c.defaultExpireInterval
	} else if expireTime < 0 {
		expireTime = c.maxExpireTime
	}
	itm := &item{
		expireTime, val,
	}
	if _,ok := c.items[key] ; ok {
		c.items[key].Value = val
		c.updateList(c.items[key])
	} else {
		c.items[key] =c.list.PushFront(itm)
	}
}