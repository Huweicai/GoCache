package main

import (
	"github.com/Huweicai/GoCache/cache"
	"strconv"
	"time"
)

func main() {
	//init := time.Now().Unix()
	cache, _ := cache.NewCache(10, 3, func(key string, value interface{}) {
		println(key + "out")
	})
	cache.Set("a", "123")
	cache.Set("b", "234")
	for i := 0; i < 1000; i++ {
		cache.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
	time.Sleep(time.Duration(4) * time.Second)
	x := cache.Get("999")
	x = cache.Get("998")
	x = cache.Get("997")
	x = cache.Get("993")
	x = cache.Get("995")
	x = cache.Get("994")
	x = cache.Get("996")
	y := cache.Get("a")
	if x != nil {
		println(x, y)
	}
}
