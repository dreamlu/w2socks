package main

import (
	"fyne.io/systray"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui"
	"github.com/dreamlu/w2socks/client/gui/global"
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
	global.Mmin = systray.AddMenuItem("最小化", "mini")
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	mUrl.Hide()
	//systray.AddSeparator() // 分隔线
	for {
		select {
		case <-mUrl.ClickedCh:
			global.Mmin.Show()
			mUrl.Hide()
			global.G.Show()
			//o <- 0
		case <-mQuit.ClickedCh:
			systray.Quit()
			global.G.Close()
			return
		case <-global.Mmin.ClickedCh:
			mUrl.Show()
			global.Mmin.Hide()
			global.G.Hide()
		}
	}
}
