package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/systray"
	"github.com/dreamlu/w2socks/client/data"
)

var (
	// G 全局
	// 主界面
	G Window

	// CONFIG 全局
	// 选中/右键文本赋值
	CONFIG CONGIG

	// Mmin 系统托盘
	Mmin *systray.MenuItem
)

type Window struct {
	fyne.Window
	Refresh chan byte
}

type CONGIG struct {
	data.Config
	Index int
}
