package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/gui/connect"
	"image/color"
)

type SelectClickText struct {
	widget.Label
	data.Config
}

func NewSelectClickText(content string, c color.Color) *SelectClickText {
	return &SelectClickText{Label: widget.Label{
		Text: content,
	}}
}

func (c *SelectClickText) Tapped(e *fyne.PointEvent) {
	fmt.Println("left click at", e)
	connect.Name = c.Name
	connect.SerIpAddr = c.ServerIpAddr
	connect.LocalPort = c.LocalPort
}

func (c *SelectClickText) TappedSecondary(e *fyne.PointEvent) {
	fmt.Println("right click at", e)
	connect.Name = c.Name
	connect.SerIpAddr = c.ServerIpAddr
	connect.LocalPort = c.LocalPort
	var menu = fyne.NewMenu("", connect.AddItem(), connect.EditItem(), connect.DelItem())
	widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(c), e.AbsolutePosition)
}

func (c *SelectClickText) DoubleTapped(e *fyne.PointEvent) {
	fmt.Println("double click at", e)
}
