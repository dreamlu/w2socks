package toolbar

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// 工具栏
func Toolbar() fyne.CanvasObject {
	return widget.NewToolbar(
		Conn(),
		Copy(),
		Share(),
	)
}
