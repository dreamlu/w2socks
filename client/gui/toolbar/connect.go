package toolbar

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/gui/connect"
	"github.com/dreamlu/w2socks/client/gui/window"
)

func Conn() widget.ToolbarItem {
	return widget.NewToolbarAction(theme.ConfirmIcon(), func() {
		window.Connect(connect.SerIpAddr, connect.LocalPort)
	})
}
