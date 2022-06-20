package copy

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/util/notify"
)

func Copy() {
	//复制到剪切板
	clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()
	if &global.CONFIG != nil {
		conf := global.CONFIG
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
