package main

import (
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/getlantern/systray"
)

// 开发环境: ubuntu
// 安装依赖: sudo apt-get install libgl1-mesa-dev xorg-dev libgtk-3-dev libappindicator3-dev -y

// 运行方式:
// 1.命令行
// 2.GUI
func main() {
	global.G = global.Window{
		Window:  gui.Gui(),
		Refresh: make(chan byte),
	}
	global.G.Show()
	go systray.Run(onReady, nil)
	global.G.ShowAndRun()
}

// 驻后台
func onReady() {
	systray.SetTemplateIcon(data.LogoData, data.LogoData)
	//systray.SetTitle("w2socks")
	systray.SetTooltip("w2socks")
	// 托盘菜单
	mUrl := systray.AddMenuItem("恢复", "my home")
	mMin := systray.AddMenuItem("最小化", "mini")
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	mUrl.Hide()
	//systray.AddSeparator() // 分隔线
	for {
		select {
		case <-mUrl.ClickedCh:
			mMin.Show()
			mUrl.Hide()
			global.G.Show()
			//o <- 0
		case <-mQuit.ClickedCh:
			systray.Quit()
			return
		case <-mMin.ClickedCh:
			mUrl.Show()
			mMin.Hide()
			global.G.Hide()
		}
	}
}
