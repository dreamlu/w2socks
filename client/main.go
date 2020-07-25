package main

import (
	"github.com/dreamlu/w2sockets/client/handle"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/url"
	"sync"
)

func main() {
	// 启动监听
	LocalListenAddr, err := net.ResolveTCPAddr("tcp", ":8083")
	if err != nil {
		log.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", LocalListenAddr)
	if err != nil {
		log.Panic(err)
	}
	defer l.Close()

	// ws client
	//httpClient(l)

	for {
		conn, _ := l.Accept()
		go socks2ws(conn.(*net.TCPConn))
	}
}

// client request ws server

func httpClient(l *net.TCPListener) {

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8082", Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 监听对于本机的请求
	for {
		// 发送数据
		client, err := l.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}

		// localConn被关闭时直接清除所有数据 不管没有发送的数据
		_ = client.SetLinger(0)
		// 1.处理本机的代理请求
		// 2.与server建立websocket连接
		// 3.异步发送数据段
		go func() {
			defer client.Close()
			go func() {
				err := handle.Receive(client, c)
				if err != nil {
					// 在 copy 的过程中可能会存在网络超时等 error 被 return，只要有一个发生了错误就退出本次工作
					client.Close()
				}
			}()
			handle.Send(client, c)
		}()
	}
}

// 1.处理本机的代理请求
// 2.与server建立websocket连接
func socks2ws(socks *net.TCPConn) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8082", Path: "/"}
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
