package window

import (
	"fmt"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/util/ip"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

// 进行连接
func Connect(serverIpAddr, localIpPort string) bool {
	log.Println("Server ip address: " + serverIpAddr)

	if !CheckEntry(serverIpAddr, localIpPort) {
		return false
	}

	// 退出旧携程
	if !Disconnect(localIpPort) {
		fmt.Print("err dis conn")
	}

	go core.Core(serverIpAddr, localIpPort)
	//core.Online++

	// 系统通知
	notify.SysNotify("w2socks", "success to connect "+serverIpAddr)
	return true
}

// 取消连接
func Disconnect(localPort string) bool {
	// 退出旧携程
	for k, v := range core.Ws {
		if k == localPort {
			v.CancelFunc()
			delete(core.Ws, localPort)
			return true
		}
	}
	return false
}

func CheckEntry(serverIpAddr, localIpPort string) bool {
	// ip地址是否正确
	msg, ok := ip.Check(serverIpAddr)
	if !ok {
		notify.SysNotify("warn!!", msg)
		return false
	}

	//本地端口是否正确
	if !ip.CheckPort(localIpPort) {
		notify.SysNotify("warn!!", "Incorrect local port")
		return false
	}
	return true
}
