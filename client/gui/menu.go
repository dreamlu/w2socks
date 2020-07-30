package gui

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/gui/connect"
)

// 主窗口
func MainMenu() *fyne.MainMenu {
	// file
	newItem := fyne.NewMenuItem("Import", func() {
	})
	exportItem := fyne.NewMenuItem("Export", func() {
	})

	// Connect
	editItom := fyne.NewMenuItem("Edit", func() {

	})
	delItom := fyne.NewMenuItem("Delete", func() {

	})

	helpMenu := fyne.NewMenu("Help", fyne.NewMenuItem("About", func() {
		// TODO 超链接
	}))

	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first mainMenu
		fyne.NewMenu("File", newItem, exportItem),
		fyne.NewMenu("Connect", connect.AddItem(), editItom, delItom),
		helpMenu,
		fyne.NewMenu("Back", fyne.NewMenuItem("back", func() {
			G.Hide()
		})),
	)

	return mainMenu
}
