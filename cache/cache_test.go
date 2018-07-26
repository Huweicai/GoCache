package cache

import (
	"strconv"
	"testing"
	"time"
)

func TestCache_GetestMap(t *testing.T) {

}

func TestNewCache(t *testing.T) {
	//init := time.Now().Unix()
	cache, _ := NewCache(10, 3, func(key string, value interface{}) {
	})
	cache.Set("a", "123")
	cache.Set("b", "234")
	for i := 0; i < 1000; i++ {
		cache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
	time.Sleep(400)
	x := cache.Get("999")
	x = cache.Get("998")
	x = cache.Get("997")
	x = cache.Get("993")
	x = cache.Get("995")
	x = cache.Get("994")
	x = cache.Get("996")
	y := cache.Get("a")
	t.Log(cache.Get("a"))
	println(x, y)

}
