package server

import (
	"log"
	"net/http"
)

func httpServerInit() {
	http.HandleFunc("/GET", httpHandler)
	err := http.ListenAndServe("127.0.0.1:6789", nil)
	if err != nil {
		log.Fatal("HTTP server started failed")
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {

}
