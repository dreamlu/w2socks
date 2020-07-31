package gui

import (
	"fmt"
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

	// 连接信息列表
	infos := widget.NewVBox()
	//infos.Resize(fyne.NewSize(250, 200))

	conf := data.GetConfig()
	for _, v := range conf {
		infos.Children = append(infos.Children, canvas.NewLine(color.White))
		text := NewSelectClickText(fmt.Sprintf("%s [%s,%s]", v.Name, v.ServerIpAddr, v.LocalPort), *v)
		infos.Children = append(infos.Children, text)
	}
	infos.Children = append(infos.Children, canvas.NewLine(color.White))

	// 纵向列表
	list := widget.NewScrollContainer(infos)
	list.Resize(fyne.NewSize(250, 200))

	// 主菜单
	addWindow.SetMainMenu(MainMenu())
	// 布局
	top := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), Toolbar())
	bom := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), list.Content)
	oth := fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	addWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), top, bom, oth))
	return addWindow
}
