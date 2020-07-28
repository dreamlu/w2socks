package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/url"
	"sync"
)

func main() {
	app := app.New()
	w := app.NewWindow("w2socks")
	w.Resize(fyne.NewSize(280, 300))

	comSize := fyne.NewSize(100, 20)

	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("ip:port                       ")
	ipEntry.Resize(comSize)

	lpEntry := widget.NewEntry()
	lpEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("ip addr:", ipEntry),
		widget.NewFormItem("local port:", lpEntry),
	)

	form.CancelText = "cancel"
	form.SubmitText = "submit"
	form.OnCancel = func() {
		log.Println("取消")
	}
	form.OnSubmit = func() {
		log.Println("提交")
		//fyne.CurrentApp().SendNotification(&fyne.Notification{
		//	Title:   "form submit",
		//	Content: ipEntry.Text,
		//})
		log.Println(ipEntry.Text)
		// 退出旧携程
		if online > 0 {
			quit <- 1
		}
		go client(ipEntry.Text, lpEntry.Text)
		online++
	}

	content := widget.NewVBox(
		widget.NewHBox(
			//widget.NewLabel("ip addr:"),
			form,
		),
	)
	w.SetContent(content)
	w.ShowAndRun()

}

var (
	online int
	quit   = make(chan byte)
)

// client
func client(ipAddr, localPort string) {
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
		case <-quit:
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
	u := url.URL{Scheme: "ws", Host: ipAddr, Path: "/"} // 149.28.34.65:8018
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
