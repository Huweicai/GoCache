package cache

import "time"

/*
	function set add or update an item to Cache
	1 check expireTime
	2 check if exists
	3 update or insert
	item expires after expireTime when it's set positive,
	but if expireTime equals zero , it uses default expireTime,
	and item never expires while it's negative.
*/
func (c *Cache)set(key string , val interface{} , expireInterval int64)  {
	var expireTime int64
	now := time.Now().Unix()
	switch  {
	case expireInterval>0:
		expireTime = now + expireInterval
	case expireInterval==0:
		expireTime = now + c.defaultExpireInterval
	case expireInterval <0:
		expireTime = c.maxExpireTime
	}
	itm := &item{
		expireTime,  key,val,
	}
	//it's not necessary to lock the whole function
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.size() >= c.limit {
		c.RemoveOldest()
	}
	if _,ok := c.items[key] ; ok {
		c.items[key].Value = val
		c.updateList(c.items[key])
	} else {
		c.items[key] =c.list.PushFront(itm)
	}
}

/*
	Default set
 */
func (c *Cache)Set(key string , val interface{})  {
	c.set(key , val , 0)
}

/*
	Set with eXpire interval
 */
func (c *Cache)SetX(key string , val interface{} , expireInterval int64 )  {
	c.set(key , val , expireInterval)
}




