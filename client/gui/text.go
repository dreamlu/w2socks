package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/connect"
)

type SelectClickText struct {
	widget.Label
	data.Config
}

func NewSelectClickText(content string, conf data.Config) *SelectClickText {
	return &SelectClickText{
		Label: widget.Label{
			Text: content,
		},
		Config: conf,
	}
}

func (c *SelectClickText) Tapped(e *fyne.PointEvent) {
	connect.Name = c.Name
	connect.SerIpAddr = c.ServerIpAddr
	connect.LocalPort = c.LocalPort
}

func (c *SelectClickText) TappedSecondary(e *fyne.PointEvent) {
	connect.Name = c.Name
	connect.SerIpAddr = c.ServerIpAddr
	connect.LocalPort = c.LocalPort
	var menu = fyne.NewMenu("", connect.AddItem(), connect.EditItem(), connect.DelItem())
	widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(c), e.AbsolutePosition)
}

//func (c *SelectClickText) DoubleTapped(e *fyne.PointEvent) {
//	fmt.Println("double click at", e)
//}
