package server

import "code.byted.org/gopkg/logs"

func InitTcpServer(port string)  {
	logs.Info("tcp server started at port : %s" , port)
}