package toolbar

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func Copy() widget.ToolbarItem {
	return widget.NewToolbarAction(theme.ContentCopyIcon(), func() {

	})
}
