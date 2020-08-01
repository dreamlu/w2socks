package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/window"
)

// 编辑逻辑
func EditItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Edit", func() {
		w := window.Window(&CONFIG, false)
		w.Show()
	})
}