package toolbar

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Share() widget.ToolbarItem {
	return widget.NewToolbarAction(theme.MailSendIcon(), func() {
	})
}
