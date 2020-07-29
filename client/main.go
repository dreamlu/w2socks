package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/core"
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
	app := app.New()
	//app.SetIcon(data.Logo())
	w := app.NewWindow("w2socks")
	w.Resize(fyne.NewSize(280, 300))

	comSize := fyne.NewSize(100, 20)

	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("ip:port")
	ipEntry.Resize(comSize)

	lpEntry := widget.NewEntry()
	lpEntry.SetPlaceHolder("port")
	lpEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("ip addr:", ipEntry),
		widget.NewFormItem("local port:", lpEntry),
	)

	form.CancelText = "off"
	form.SubmitText = "ok"
	form.OnCancel = func() {
		log.Println("取消")
	}
	form.OnSubmit = func() {
		log.Println("提交")
		//fyne.CurrentApp().SendNotification(&fyne.Notification{
		//	Title:   "form submit",
		//	Content: ipEntry.Text,
		//})
		log.Println(ipEntry.Text)
		// 退出旧携程
		if core.Online > 0 {
			core.Quit <- 1
		}
		go core.Core(ipEntry.Text, lpEntry.Text)
		core.Online++
	}

	content := widget.NewVBox(
		widget.NewVBox(
			//widget.NewLabel("ip addr:"),
			form,
		),
	)
	w.SetContent(content)
	w.ShowAndRun()
}
