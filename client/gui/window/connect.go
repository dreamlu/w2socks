package window

import (
	"fmt"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/util/ip"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

// 进行连接
func Connect(c core.W2Config) bool {
	log.Println("Server ip address: " + c.ServerIpAddr)

	if !CheckEntry(c.ServerIpAddr, c.LocalPort) {
		return false
	}

	// 退出旧携程
	if !Disconnect(c.String()) {
		fmt.Print("err dis conn")
	}

	go core.Core(&core.W2Config{
		ServerIpAddr: c.ServerIpAddr,
		LocalPort:    c.LocalPort,
	})
	//core.Online++

	// 系统通知
	notify.SysNotify("w2socks", "success to connect "+c.ServerIpAddr)
	global.G.Refresh <- 1
	return true
}

// 取消连接
func Disconnect(key string) bool {
	// 退出旧携程
	for k, v := range core.Ws {
		if k == key {
			v.CancelFunc()
			delete(core.Ws, key)
			global.G.Refresh <- 1
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
