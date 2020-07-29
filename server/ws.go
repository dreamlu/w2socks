package main

import (
	"github.com/dreamlu/w2socks/server/handle"
	"github.com/gorilla/websocket"
	"net/http"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 消息读取
// 开启不同进程代表对应的客户端通信
func WsHander(ws *websocket.Conn) {

	defer ws.Close()
	//消息读取,每个客户端数据
	for {
		handle.Handle(ws)
	}
}
