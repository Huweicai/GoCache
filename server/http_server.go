package server

import (
	"log"
	"net/http"
	"strings"
)

func InitHttpServer(port string) {
	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe("127.0.0.1:" + port, nil)
	if err != nil {
		log.Fatal("HTTP server started failed " + err.Error())
	}
	log.Printf("HTTP server started at port : %s" , port)
}

/*
Restful http handler
 */
func httpHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path , "/")
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodDelete:
	case http.MethodPut:

	}
	log.Println("httpHandler in key :" + key)
}


