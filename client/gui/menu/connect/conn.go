package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/window"
)

// 全局
// 选中/右键文本赋值
var (
	CONFIG data.Config
)

// 连接逻辑
func ConnItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Connect", func() {
		window.Connect(CONFIG.ServerIpAddr, CONFIG.LocalPort)
	})
}