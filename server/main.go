package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpServer)
	err := http.ListenAndServe(":8018", nil)
	if err != nil {
		log.Println("ListenAndServe error: ", err)
	}
}

func httpServer(w http.ResponseWriter, r *http.Request) {

	ws, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("httpServer error: ", err)
		return
	}
	WsHandler(ws)
}
