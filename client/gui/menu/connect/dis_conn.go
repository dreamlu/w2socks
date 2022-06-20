package connect

import (
	"fyne.io/fyne/v2"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/gui/window"
)

// 取消连接逻辑
func DisConnItem() *fyne.MenuItem {
	return fyne.NewMenuItem("DisConn", func() {
		window.Disconnect(global.CONFIG.W2Config.String())
	})
}
