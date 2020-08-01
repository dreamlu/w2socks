package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/gui/window"
)

// 连接逻辑
func ConnItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Connect", func() {
		window.Connect(global.CONFIG.ServerIpAddr, global.CONFIG.LocalPort)
	})
}
