package window

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

// 通用window
// 编辑/添加连接窗体
func Window(ipAddr, port string) fyne.Window {
	w := fyne.CurrentApp().NewWindow("connect content")
	w.Resize(fyne.NewSize(280, 300))
	comSize := fyne.NewSize(100, 20)

	// m名字
	nameEntry := widget.NewEntry()
	if ipAddr == "" {
		nameEntry.SetPlaceHolder("name")
	} else {
		nameEntry.Text = ipAddr
	}
	nameEntry.Resize(comSize)

	// 服务端ip和端口
	serverEntry := widget.NewEntry()
	if ipAddr == "" {
		serverEntry.SetPlaceHolder("ip:port")
	} else {
		serverEntry.Text = ipAddr
	}
	serverEntry.Resize(comSize)

	// 本地端口号
	localPortEntry := widget.NewEntry()
	if port == "" {
		localPortEntry.SetPlaceHolder("port")
	} else {
		localPortEntry.Text = port
	}
	localPortEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("name:", nameEntry),
		widget.NewFormItem("server:", serverEntry),
		widget.NewFormItem("local:", localPortEntry),
	)

	form.CancelText = "cancel"
	form.SubmitText = "save"

	// 取消操作
	form.OnCancel = func() {
		//if Disconnect() {
		//	notify.SysNotify("notify", "server is disconnected")
		//}
		w.Close()
	}

	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		b := CheckEntry(serverEntry.Text, localPortEntry.Text)
		if !b {
			return
		}
		// TODO 编辑也是添加问题
		// TODO 删除问题
		conf := data.Config{}
		conf.Name = nameEntry.Text
		conf.ServerIpAddr = serverEntry.Text
		conf.LocalPort = localPortEntry.Text
		err := data.SaveConfig(conf)
		if err != nil {
			notify.SysNotify("warn!!", err.Error())
			return
		}
		notify.SysNotify("info", "连接信息已存入")
		w.Close()
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
