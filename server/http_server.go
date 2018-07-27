package server

import (
	"log"
	"net/http"
)

func InitHttpServer(port string) {
	http.HandleFunc("/GET", httpHandler)
	err := http.ListenAndServe("127.0.0.1:" + port, nil)
	if err != nil {
		log.Fatal("HTTP server started failed")
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

}


