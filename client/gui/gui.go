package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/cuscom/text"
	"github.com/dreamlu/w2socks/client/gui/menu"
	"github.com/dreamlu/w2socks/client/gui/toolbar"
	"image/color"
)

var majorWindow fyne.Window

// 主窗体
func Gui() fyne.Window {
	// 主程序
	majorApp := app.New()
	// light theme
	majorApp.Settings().SetTheme(theme.LightTheme())
	// logo
	majorApp.SetIcon(data.Logo())
	majorWindow = majorApp.NewWindow("w2socks")
	size := fyne.NewSize(280, 300)
	majorWindow.Resize(size)

	// 主菜单
	majorWindow.SetMainMenu(menu.MainMenu())

	// 布局
	majorWindow.SetContent(Content())
	return majorWindow
}

// 列表核心
func mainList() []fyne.CanvasObject {
	var items []fyne.CanvasObject

	// 获取本地配置并加载到容器
	conf := data.GetConfig()
	for _, v := range conf {
		item := widget.NewVBox(
			text.NewSelectClickText(fmt.Sprintf("%s\n%s", v.Name, v.ServerIpAddr), *v),
			canvas.NewLine(color.Black),
		)
		items = append(items, item)
	}
	return items
}

// 刷新主界面内容
func Content() fyne.CanvasObject {
	top := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), toolbar.Toolbar())
	list := []fyne.CanvasObject{top}
	list = append(list, mainList()...)
	vert := widget.NewVScrollContainer(widget.NewVBox(list...))
	return vert
}
