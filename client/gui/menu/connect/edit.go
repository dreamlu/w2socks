package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/gui/window"
	"github.com/dreamlu/w2socks/client/util/notify"
)

// 编辑逻辑
func EditItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Edit", func() {
		if &global.CONFIG == nil {
			notify.SysNotify("warn!!", "Edit content not selected")
			return
		}
		window.OpenWindow(&global.CONFIG.Config, false)
		//global.G.Modal <- 1
	})
}
