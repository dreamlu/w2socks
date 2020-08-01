package toolbar

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/gui/window"
	"github.com/dreamlu/w2socks/client/util/notify"
)

func Conn() widget.ToolbarItem {
	return widget.NewToolbarAction(theme.ConfirmIcon(), func() {
		if &global.CONFIG != nil {
			window.Connect(global.CONFIG.ServerIpAddr, global.CONFIG.LocalPort)
		} else {
			notify.SysNotify("warn!!", "No content selected")
		}
	})
}
