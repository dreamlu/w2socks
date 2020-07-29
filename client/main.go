package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/util/ip"
	"log"
)

// 运行方式:
// 1.命令行
// 2.GUI
func main() {
	window()
}

// window
func window() {
	// 主程序
	majorApp := app.New()

	// logo
	majorApp.SetIcon(data.Logo())
	// 主窗体
	mainWindow := majorApp.NewWindow("w2socks")
	mainWindow.Resize(fyne.NewSize(280, 300))

	comSize := fyne.NewSize(100, 20)

	// 服务端ip和端口
	serviceIpPortEntry := widget.NewEntry()
	serviceIpPortEntry.SetPlaceHolder("ip:port")
	serviceIpPortEntry.Resize(comSize)

	// 本地端口号
	localPortEntry := widget.NewEntry()
	localPortEntry.SetPlaceHolder("port")
	localPortEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("server:", serviceIpPortEntry),
		widget.NewFormItem("local:", localPortEntry),
	)

	form.CancelText = "disconnect"
	form.SubmitText = "connect"
	// 取消操作
	form.OnCancel = func() {
		SysNotify("notify", "server is disconnected")
		log.Println("取消")
	}
	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		ipAddr := serviceIpPortEntry.Text
		log.Println("用户输入: " + ipAddr)

		// ip地址是否正确
		msg, ok := ip.Check(ipAddr)
		if !ok {
			SysNotify("warn!!", msg)
			return
		}

		//本地端口是否正确
		if !ip.CheckPort(localPortEntry.Text) {
			SysNotify("warn!!", "Incorrect local port")
			return
		}

		// 退出旧携程
		if core.Online > 0 {
			core.Quit <- 1
		}
		go core.Core(ipAddr, localPortEntry.Text)
		core.Online++
		// 系统通知
		SysNotify("w2socks", "success to connect "+ipAddr)
	}

	// 窗体
	content := widget.NewVBox(
		widget.NewVBox(
			// 输入服务端的ip地址和端口 以及本地的端口
			widget.NewLabel("Please Enter:"),
			form,
		),
	)
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
