package connect

import (
	"fyne.io/fyne/v2"
	"github.com/dreamlu/w2socks/client/gui/window"
)

// 添加逻辑
func AddItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Add", func() {
		window.OpenWindow(nil, true)
		//w.Show()
	})
}
