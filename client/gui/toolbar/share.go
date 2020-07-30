package toolbar

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func Share() widget.ToolbarItem {
	return widget.NewToolbarAction(theme.MailSendIcon(), func() {
	})
}
