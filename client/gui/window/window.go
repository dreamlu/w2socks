package window

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

// 通用window
// 编辑/添加连接窗体
func Window() fyne.Window {
	w := fyne.CurrentApp().NewWindow("connect content")
	w.Resize(fyne.NewSize(280, 300))
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
		if Disconnect() {
			notify.SysNotify("notify", "server is disconnected")
		}
		w.Close()
	}

	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		b := Connect(serverEntry.Text, localPortEntry.Text)
		if !b {
			return
		}
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
	w.SetContent(content)
	return w
}
