package toolbar

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	copy2 "github.com/dreamlu/w2socks/client/util/copy"
)

func Copy() widget.ToolbarItem {
	return widget.NewToolbarAction(theme.ContentCopyIcon(), copy2.Copy)
}
