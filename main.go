package main

import (
	"io/ioutil"
	"code.byted.org/gopkg/pkg/log"
	"encoding/json"
	"GoCache/conf"
	"GoCache/server"
	"strconv"
)

const (
	defaultHTTPPort = "6789"
	defaultTCPPort = "6788"
)

func main() {

	confFile,err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		log.Error("load config failed" + err.Error())
		return
	}
	var conf  = new(config.Conf)
	json.Unmarshal(confFile , &conf)
	httpEnable , _ := strconv.ParseBool(conf.Http.Enable)
	tcpEnable , _ := strconv.ParseBool(conf.Tcp.Enable)

	if  httpEnable {
		port := conf.Http.Port
		if port == "" {
			port = defaultHTTPPort
		}
		server.InitHttpServer(port)
	}

	if tcpEnable {
		port := conf.Tcp.Port
		if port == "" {
			port = defaultTCPPort
		}
		server.InitTcpServer(port)
	}
}
