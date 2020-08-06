package global

import (
	"fyne.io/fyne"
	"github.com/dreamlu/w2socks/client/data"
)

var (
	// 全局
	// 主界面
	G Window

	// 全局
	// 选中/右键文本赋值
	CONFIG CONGIG
)

type Window struct {
	fyne.Window
	Refresh chan byte
}

type CONGIG struct {
	data.Config
	Index int
}
