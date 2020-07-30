package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/ctrl"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui"
	"github.com/getlantern/systray"
	"log"
)

// 开发环境: ubuntu
// 安装依赖: sudo apt-get install libgl1-mesa-dev xorg-dev libgtk-3-dev libappindicator3-dev -y

var (
	w        fyne.Window
	majorApp fyne.App
)

// 运行方式:
// 1.命令行
// 2.GUI
func main() {
	// 主程序
	majorApp = app.New()

	// logo
	majorApp.SetIcon(data.Logo())

	w1 := mainForm()
	w1.Show()

	w = window()
	go systray.Run(onReady, nil)
	//w.SetOnClosed(func() {
	//	w.Hide()
	//	w.Show()
	//	systray.Run(onReady, nil)
	//})
	w.ShowAndRun()
}

// 驻后台
func onReady() {
	systray.SetTemplateIcon(data.LogoData, data.LogoData)
	//systray.SetTitle("w2socks")
	systray.SetTooltip("w2socks")
	// 托盘菜单
	mUrl := systray.AddMenuItem("恢复", "my home")
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	//systray.AddSeparator() // 分隔线
	for {
		select {
		case <-mUrl.ClickedCh:
			w.Show()
			//o <- 0
		case <-mQuit.ClickedCh:
			//o <- 1
			return
		}
	}
}

// 主窗体
func mainForm() fyne.Window {
	addWindow := majorApp.NewWindow("w2socks")

	addWindow.Resize(fyne.NewSize(280, 300))

	// 纵向列表
	infos := widget.NewVBox()

	conf := data.GetConfig()
	for _, v := range conf {
		infos.Children = append(infos.Children, widget.NewLabel(v.Name))
	}

	list := widget.NewVScrollContainer(infos)
	list.Resize(fyne.NewSize(130, 200))

	addWindow.SetMainMenu(gui.MainMenu())
	addWindow.SetContent(widget.NewVBox(infos))
	return addWindow
}

// 连接窗体
func window() fyne.Window {
	// 主窗体
	mainWindow := majorApp.NewWindow("add new conn info")
	mainWindow.Resize(fyne.NewSize(280, 300))
	comSize := fyne.NewSize(100, 20)

	// 服务端ip和端口
	serverEntry := widget.NewEntry()
	serverEntry.SetPlaceHolder("ip:port")
	serverEntry.Resize(comSize)

	// 本地端口号
	localPortEntry := widget.NewEntry()
	localPortEntry.SetPlaceHolder("port")
	localPortEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("server:", serverEntry),
		widget.NewFormItem("local:", localPortEntry),
	)

	form.CancelText = "disconnect"
	form.SubmitText = "connect"

	// 取消操作
	form.OnCancel = func() {
		ctrl.Disconnect()
	}

	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		ctrl.Connect(serverEntry.Text, localPortEntry.Text)
		w.Hide()
	}
	//form.OnSubmit = func() {
	//	log.Println("提交")
	//	ctrl.ConnectAndCall(serverEntry.Text, localPortEntry.Text, w.Hide)
	//}

	// 窗体
	content := widget.NewVBox(
		widget.NewVBox(
			// 输入服务端的ip地址和端口 以及本地的端口
			widget.NewLabel("Please Enter:"),
			form,
		),
	)
	mainWindow.SetContent(content)
	//mainWindow.ShowAndRun()
	return mainWindow
}
