package text

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/global"
	"github.com/dreamlu/w2socks/client/gui/menu/connect"
)

// 可以点击和右选中的文本
type SelectClickText struct {
	*widget.Label
	data.Config
	index int
	//BaseRenderer
	//layout fyne.Layout
}

func NewSelectClickText(content string, conf data.Config, index int) *SelectClickText {
	lb := widget.NewLabel(content)
	//lb.Resize(fyne.NewSize(40, 80))
	t := &SelectClickText{
		Label:  lb,
		Config: conf,
		index:  index,
	}

	// 判断是否连接
	for _, v := range core.Ws {
		if v.Value("localPort").(string) == t.LocalPort {
			lb.TextStyle = fyne.TextStyle{
				Bold:   true,
				Italic: true,
			}
			break
		}
	}
	return t
}

func (c *SelectClickText) Tapped(e *fyne.PointEvent) {
	global.CONFIG = global.CONGIG{
		Config: c.Config,
		Index:  c.index,
	}
}

func (c *SelectClickText) TappedSecondary(e *fyne.PointEvent) {
	global.CONFIG = global.CONGIG{
		Config: c.Config,
		Index:  c.index,
	}
	var menu *fyne.Menu
	menu = fyne.NewMenu("", connect.ConnItem(), connect.AddItem(), connect.EditItem(), connect.DelItem())
	// 判断是否连接
	for _, v := range core.Ws {
		if v.Value("localPort").(string) == c.Config.LocalPort {
			menu = fyne.NewMenu("", connect.DisConnItem(), connect.AddItem(), connect.EditItem(), connect.DelItem())
			break
		}
	}

	widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(c), e.AbsolutePosition)
}

//func (c *SelectClickText) DoubleTapped(e *fyne.PointEvent) {
//	fmt.Println("double click at", e)
//}

//func (c *SelectClickText) BackgroundColor() color.Color {
//	return color.RGBA{R: 0, G: 150, B: 200, A: 255}
//}

//func (c *SelectClickText) CreateRenderer() fyne.WidgetRenderer {
//	//c.ExtendBaseWidget(c)
//	return c
//}

//func (c *SelectClickText) Layout(size fyne.Size) {
//	c.layout.Layout(c.Objects(), size)
//}
