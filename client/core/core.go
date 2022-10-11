package core

import (
	"context"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	Ws = map[string]W2socks{}
)

type W2Config struct {
	ServerIpAddr string `json:"server_ip_addr"`
	LocalPort    string `json:"local_port"`
}

func (w *W2Config) String() string {
	return w.ServerIpAddr + ";" + w.LocalPort
}

type W2socks struct {
	context.Context
	context.CancelFunc
}

// W2socket websocket
type W2socket struct {
	*websocket.Conn
	port string
}

// Core client
func Core(wc *W2Config) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "localPort", wc.LocalPort)

	Ws[wc.String()] = W2socks{
		Context:    ctx,
		CancelFunc: cancel,
	}
	//context.WithDeadline()
	//go telnetLocal(wc.LocalPort)
	listen(wc.ServerIpAddr, wc.LocalPort)
}

// client listen
func listen(ipAddr, localPort string) {
	// 启动监听
	LocalListenAddr, err := net.ResolveTCPAddr("tcp", ":"+localPort)
	if err != nil {
		log.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", LocalListenAddr)
	if err != nil {
		log.Println("listen tcp error: ", err)
	}
	defer l.Close()

	// 原理: 原先的每一次的http连接用websocket连接替代
	for {
		// 预先建立websocket通道
		ws := websockets(ipAddr)
		// 新的请求
		conn, _ := l.AcceptTCP()
		if conn == nil {
			continue
		}
		//conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		go socks2ws(conn, ws)
	}
}

// 1.处理本机的代理请求
// 2.与server进行数据通信
func socks2ws(socks *net.TCPConn, ws *websocket.Conn) {

	defer func() {
		log.Println("ws该次连接关闭")
		ws.Close()
	}()
	var wg sync.WaitGroup
	ioCopy := func(dst io.Writer, src io.Reader) {
		log.Println("数据通信")
		defer wg.Done()
		io.Copy(dst, src)
	}
	wg.Add(2)
	go ioCopy(ws.UnderlyingConn(), socks)
	go ioCopy(socks, ws.UnderlyingConn())
	wg.Wait()
}

// 建立websocket连接
func websockets(ipAddr string) *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: ipAddr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("dial error:", err)
		log.Println("trying after 3s")
		time.Sleep(3 * time.Second)
		return websockets(ipAddr)
	}
	return ws
}

// telnet local
func telnetLocal(localPort string) {
	var (
		cmd   *exec.Cmd
		bytes []byte
	)
	for range time.Tick(time.Millisecond * 500) {
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("nc", "-vz", "-w", "2", "127.0.0.1", localPort)
			//读取所有输出
			bytes, _ = cmd.Output()
			if !strings.Contains(string(bytes), "succeeded") {
				return
			}
		default:
			cmd = exec.Command("telnet", "127.0.0.1", localPort)
			//读取所有输出
			bytes, _ = cmd.Output()
			if !strings.Contains(string(bytes), "Connected") {
				return
			}
		}
		log.Println(string(bytes))
	}
}
