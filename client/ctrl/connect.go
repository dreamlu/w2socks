package ctrl

import (
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/util/ip"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

// 进行连接
func Connect(serverIpAddr, localIpPort string) {
	log.Println("Server ip address: " + serverIpAddr)

	// ip地址是否正确
	msg, ok := ip.Check(serverIpAddr)
	if !ok {
		notify.SysNotify("warn!!", msg)
		return
	}

	//本地端口是否正确
	if !ip.CheckPort(localIpPort) {
		notify.SysNotify("warn!!", "Incorrect local port")
		return
	}

	// 退出旧携程
	if core.Online > 0 {
		core.Quit <- 1
	}

	go core.Core(serverIpAddr, localIpPort)
	core.Online++

	// 系统通知
	notify.SysNotify("w2socks", "success to connect "+serverIpAddr)
}

// 进行连接
func ConnectAndCall(serverIpAddr, localIpPort string, call func()) {
	Connect(serverIpAddr, localIpPort)
	call()
}

// 取消连接
func Disconnect() {
	// 退出旧携程
	if core.Online > 0 {
		core.Quit <- 1
		notify.SysNotify("notify", "server is disconnected")
	}
}
