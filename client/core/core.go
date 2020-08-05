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
	Wt = map[string][]*W2socket{}
)

type W2socks struct {
	context.Context
	context.CancelFunc
}

// websocket
type W2socket struct {
	*websocket.Conn
	port string
}

// close
func CloseContext(ctx context.Context) {

	localPort := ctx.Value("localPort").(string)
	for _, v := range Wt {
		for _, v2 := range v {
			if v2.port == localPort {
				v2.UnderlyingConn().Close()
			}
		}
	}
}

// client
func Core(ipAddr, localPort string) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "localPort", localPort)

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
		ws := websockets(ipAddr)
		ws.port = localPort
		Wt[localPort] = append(Wt[localPort], ws)
		conn, _ := l.Accept()
		select {
		case <-ctx.Done():
			CloseContext(ctx)
			//ws.Close()
			fmt.Printf("旧线程结束\n")
			return
		default:
			if conn == nil {
				return
			}
			go socks2ws(conn.(*net.TCPConn), ws.Conn)
		}
	}
}

// 建立websocket连接
func websockets(ipAddr string) *W2socket {
	u := url.URL{Scheme: "ws", Host: ipAddr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return &W2socket{
		Conn: ws,
	}
}

// 1.处理本机的代理请求
// 2.与server建立websocket连接
func socks2ws(socks *net.TCPConn, ws *websocket.Conn) {

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
