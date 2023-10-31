package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var p string
	flag.StringVar(&p, "p", "8018", "本地监听端口")
	flag.Parse()
	log.Printf("local listen port: %s", p)
	http.HandleFunc("/", httpServer)
	err := http.ListenAndServe(":"+p, nil)
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
