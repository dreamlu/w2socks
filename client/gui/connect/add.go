package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/window"
)

// 添加逻辑
func AddItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Add", func() {
		w := window.Window("", "")
		w.Show()
	})
}
