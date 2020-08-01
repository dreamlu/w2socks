package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/window"
	"github.com/dreamlu/w2socks/client/util/notify"
)

// 编辑逻辑
func EditItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Edit", func() {
		if &CONFIG == nil || CONFIG.ID == "" {
			notify.SysNotify("warn!!", "Edit content not selected")
			return
		}
		w := window.Window(&CONFIG, false)
		w.Show()
	})
}
