package cache

import "container/list"

func (c *cache)updateList(elmt *list.Element)  {
	c.list.MoveToFront(elmt)
}
