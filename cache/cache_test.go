package cache

import (
	"testing"
	"strconv"
)

func TestCache_GetestMap(t *testing.T) {

}

func TestNewCache(t *testing.T) {
	//init := time.Now().Unix()
	cache,_ := NewCache(10,3, func(key string, value interface{}) {
	})
	cache.Set("a" , "123" )
	cache.Set("b" , "234" )
	for i:=0; i < 1000 ; i++ {
		cache.Set(strconv.Itoa(i) , strconv.Itoa(i))
	}
	x := cache.Get("a")
	y := cache.Get("a")
	t.Log(	cache.Get("a"))
	println(x , y)

}
