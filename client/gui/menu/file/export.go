package file

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/dreamlu/w2socks/client/gui/global"
	"log"
)

// 导出逻辑
func ExportItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Export", func() {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, global.G)
				return
			}
			fileSaved(writer)
		}, global.G)
	})
}

func fileSaved(f fyne.URIWriteCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	log.Println("Save to...", f.URI())
}
