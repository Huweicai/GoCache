package main

import (
	"io/ioutil"
	"code.byted.org/gopkg/pkg/log"
	"encoding/json"
	"GoCache/conf"
)

const (
	defaultHTTPPort = 6789
)

func main() {

	confFile,err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		log.Error("load config failed" + err.Error())
		return
	}
	var conf  = new(config.Conf)
	json.Unmarshal(confFile , &conf)
	println(conf)
}
