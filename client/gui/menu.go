package gui

import (
	"fyne.io/fyne"
)

func MainMenu() *fyne.MainMenu {
	// file
	newItem := fyne.NewMenuItem("Import", func() {
	})
	exportItem := fyne.NewMenuItem("Export", func() {
	})

	// Connect
	addItom := fyne.NewMenuItem("Add", func() {

	})
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
		fyne.NewMenu("Connect", addItom, editItom, delItom),
		helpMenu,
	)

	return mainMenu
}
