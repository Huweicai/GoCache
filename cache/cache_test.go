package cache

import (
	"testing"
)

func TestCache_GetestMap(t *testing.T) {

}

func TestNewCache(t *testing.T) {
	//init := time.Now().Unix()
	cache,_ := NewCache(10,30,nil)
	cache.Set("a" , "123" , 0)
	cache.Set("b" , "234" , 0)
	t.Log(	cache.Get("a"))

}
