package connect

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/util/notify"
)

// 删除逻辑
func DelItem() *fyne.MenuItem {
	return fyne.NewMenuItem("Delete", func() {
		err := data.DeleteConfig(CONFIG.ID)
		if err != nil {
			notify.SysNotify("error!!", err.Error())
		}
	})
}
