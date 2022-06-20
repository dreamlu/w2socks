package menu

import (
	"fyne.io/fyne/v2"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/gui/menu/connect"
	"github.com/dreamlu/w2socks/client/gui/menu/file"
	"github.com/dreamlu/w2socks/client/gui/menu/help"
)

// 主窗口菜单
func MainMenu() *fyne.MainMenu {
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first mainMenu
		fyne.NewMenu("File", file.ImportItem(), file.ExportItem()),
		fyne.NewMenu("Connect", connect.AddItem(), connect.EditItem(), connect.DelItem()),
		fyne.NewMenu("Help", help.HelpItem()),
		fyne.NewMenu("Back", fyne.NewMenuItem("back", func() {
			global.Mmin.ClickedCh <- struct{}{}
		})),
	)
	return mainMenu
}
