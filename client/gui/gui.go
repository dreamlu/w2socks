package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	"image/color"
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

	conf := data.GetConfig()
	for _, v := range conf {
		infos.Children = append(infos.Children, canvas.NewLine(color.White))
		infos.Children = append(infos.Children, NewSelectClick(v.Name, color.White))
	}
	infos.Children = append(infos.Children, canvas.NewLine(color.White))

	list := widget.NewVScrollContainer(infos)
	list.Resize(fyne.NewSize(130, 200))

	addWindow.SetMainMenu(MainMenu())
	// 布局
	top := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), Toolbar())
	bom := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), infos)
	oth := fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	addWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), top, bom, oth))
	return addWindow
}
