package main

import (
	"github.com/dreamlu/w2socks/server/handle"
	"github.com/gorilla/websocket"
	"net/http"
)

// Configure the upGrader
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandler 消息读取
// 开启不同进程代表对应的客户端通信
func WsHandler(ws *websocket.Conn) {
	//消息读取,每个客户端数据
	handle.Handle(ws)
}
