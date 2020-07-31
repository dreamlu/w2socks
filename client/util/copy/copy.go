package copy

import (
	"encoding/json"
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/connect"
	"github.com/dreamlu/w2socks/client/util/notify"
)

func Copy() {
	//复制到剪切板
	clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
	if connect.Name != "" {
		conf := data.Config{
			LocalPort:    connect.LocalPort,
			Name:         connect.Name,
			ServerIpAddr: connect.SerIpAddr,
		}
		body, err := json.Marshal(conf)
		if err != nil {
			return
		}
		clipboard.SetContent(string(body))
	} else {
		// 没有选择内容
		notify.SysNotify("warn!!", "No content selected")
	}
}
