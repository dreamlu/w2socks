package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
)

var (
	G fyne.Window
)

// 主窗体
func Gui() fyne.Window {
	// 主程序
	majorApp := app.New()
	// logo
	majorApp.SetIcon(data.Logo())
	addWindow := majorApp.NewWindow("w2socks")
	addWindow.Resize(fyne.NewSize(280, 300))

	// 纵向列表
	infos := widget.NewVBox()

	//conf := data.GetConfig()
	//for _, v := range conf {
	//	infos.Children = append(infos.Children, widget.NewLabel(v.Name))
	//}

	list := widget.NewVScrollContainer(infos)
	list.Resize(fyne.NewSize(130, 200))

	addWindow.SetMainMenu(MainMenu())
	addWindow.SetContent(widget.NewVBox(infos))
	return addWindow
}
