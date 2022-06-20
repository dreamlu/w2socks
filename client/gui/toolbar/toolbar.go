package toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// 工具栏
func Toolbar() fyne.CanvasObject {
	return widget.NewToolbar(
		Conn(),
		Copy(),
		Share(),
	)
}
