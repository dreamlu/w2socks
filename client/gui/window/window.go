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
func Window(conf *data.Config, add bool) fyne.Window {
	w := fyne.CurrentApp().NewWindow("connect content")
	w.Resize(fyne.NewSize(280, 300))
	comSize := fyne.NewSize(100, 20)

	// m名字
	nameEntry := widget.NewEntry()
	// 服务端ip和端口
	serverEntry := widget.NewEntry()
	// 本地端口号
	localPortEntry := widget.NewEntry()
	if add {
		// 添加
		nameEntry.SetPlaceHolder("name")
		serverEntry.SetPlaceHolder("ip:port")
		localPortEntry.SetPlaceHolder("port")
	} else {
		nameEntry.Text = conf.Name
		serverEntry.Text = conf.ServerIpAddr
		localPortEntry.Text = conf.LocalPort
	}
	nameEntry.Resize(comSize)
	serverEntry.Resize(comSize)
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
		w.Hide()
	}

	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		b := CheckEntry(serverEntry.Text, localPortEntry.Text)
		if !b {
			return
		}
		conf := data.Config{}
		conf.Name = nameEntry.Text
		conf.ServerIpAddr = serverEntry.Text
		conf.LocalPort = localPortEntry.Text
		var err error
		if add {
			// 添加
			err = data.InsertConfig(conf)
		} else {
			// 编辑
			err = data.UpdateConfig(conf)
		}
		if err != nil {
			notify.SysNotify("warn!!", err.Error())
			return
		}
		notify.SysNotify("info", "连接信息已存入")
		w.Close()
	}

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
