package gui

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/connect"
	"github.com/dreamlu/w2socks/client/gui/file"
	"github.com/dreamlu/w2socks/client/gui/help"
)

var (
	SerIpAddr = ""
	LocalPort = ""
)

// 主窗口
func MainMenu() *fyne.MainMenu {
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first mainMenu
		fyne.NewMenu("File", file.ImportItem(), file.ExportItem()),
		fyne.NewMenu("Connect", connect.AddItem(), connect.EditItem(SerIpAddr, LocalPort), connect.DelItem()),
		fyne.NewMenu("Help", help.HelpItem()),
		fyne.NewMenu("Back", fyne.NewMenuItem("back", func() {
			G.Hide()
		})),
	)

	return mainMenu
}
