package cache

import "testing"

func TestCache_GetestMap(t *testing.T) {
	m := make(map[string]string)
	m["1"] = "1"
	_,ok := m["1"]
	_,oks := m["2"]
	println(ok)
	println(oks)
}
