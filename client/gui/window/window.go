package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
)

var W fyne.Window

// 通用window
// 编辑/添加连接窗体
func OpenWindow(conf *data.Config, add bool) fyne.Window {
	if W != nil {
		W.Close()
	}
	W = fyne.CurrentApp().NewWindow("connect content")
	W.Resize(fyne.NewSize(280, 300))
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
		W.Hide()
	}

	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		// 检查验证输入的ip地址是否有效
		b := CheckEntry(serverEntry.Text, localPortEntry.Text)
		if !b {
			return
		}
		conf := data.Config{
			Name: nameEntry.Text,
			W2Config: core.W2Config{
				ServerIpAddr: serverEntry.Text,
				LocalPort:    localPortEntry.Text,
			},
		}

		var err error
		if add {
			// 添加
			err = data.InsertConfig(conf)
		} else {
			// 编辑
			err = data.UpdateConfig(conf, global.CONFIG.Index)
		}
		if err != nil {
			notify.SysNotify("warn!!", err.Error())
			return
		}
		notify.SysNotify("info", "连接信息已存入")
		global.G.Refresh <- 1
		W.Close()
	}

	// 窗体
	content := container.NewVBox(
		container.NewVBox(
			// 输入服务端的ip地址和端口 以及本地的端口
			widget.NewLabel("Please Enter:"),
			form,
		),
	)
	W.SetContent(content)
	W.SetOnClosed(func() {
		fmt.Println("操作完成,刷新")
		global.G.Refresh <- 1
	})
	W.Show()
	return W
}

//func Modal() {
//	for {
//		if <-global.G.Modal == 1 {
//			username := widget.NewEntry()
//			password := widget.NewPasswordEntry()
//			content := widget.NewForm(widget.NewFormItem("Username", username),
//				widget.NewFormItem("Password", password))
//
//			dialog.ShowCustomConfirm("Please Enter:", "Log In", "Cancel", content, func(b bool) {
//				if !b {
//					return
//				}
//
//				log.Println("Please Authenticate", username.Text, password.Text)
//			}, global.G.Window)
//		}
//	}
//}
