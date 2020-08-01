package core

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/url"
	"sync"
)

var (
	//Online int
	Ws = map[string]W2socks{}
)

type W2socks struct {
	context.Context
	context.CancelFunc
}

// client
func Core(ipAddr, localPort string) {
	ctx, cancel := context.WithCancel(context.Background())

	Ws[localPort] = W2socks{
		Context:    ctx,
		CancelFunc: cancel,
	}
	//context.WithDeadline()
	listen(ctx, ipAddr, localPort)
}

// client listen
func listen(ctx context.Context, ipAddr, localPort string) {
	// 启动监听
	LocalListenAddr, err := net.ResolveTCPAddr("tcp", ":"+localPort)
	if err != nil {
		log.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", LocalListenAddr)
	if err != nil {
		log.Println(err)
	}
	defer l.Close()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("旧线程结束\n")
			return
		default:
			conn, _ := l.Accept()
			go socks2ws(conn.(*net.TCPConn), ipAddr)
		}
	}
}

// 1.处理本机的代理请求
// 2.与server建立websocket连接
func socks2ws(socks *net.TCPConn, ipAddr string) {
	u := url.URL{Scheme: "ws", Host: ipAddr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()

	var wg sync.WaitGroup
	ioCopy := func(dst io.Writer, src io.Reader) {
		defer wg.Done()
		io.Copy(dst, src)
	}
	wg.Add(2)
	go ioCopy(ws.UnderlyingConn(), socks)
	go ioCopy(socks, ws.UnderlyingConn())
	wg.Wait()
}
