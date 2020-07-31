package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/gui/toolbar"
)

func Toolbar() fyne.CanvasObject {
	return widget.NewToolbar(
		toolbar.Conn(),
		toolbar.Copy(),
		toolbar.Share(),
	)
}
