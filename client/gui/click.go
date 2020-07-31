package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/gui/connect"
	"image/color"
)

type SelectClickText struct {
	widget.Label
	ServerIpAddr string
	LocalPort    string
}

func NewSelectClickText(content string, c color.Color) *SelectClickText {
	return &SelectClickText{Label: widget.Label{
		Text: content,
	}}
}

func (c *SelectClickText) Tapped(e *fyne.PointEvent) {
	fmt.Println("left click at", e)
}

func (c *SelectClickText) TappedSecondary(e *fyne.PointEvent) {
	fmt.Println("right click at", e)
	var menu = fyne.NewMenu("", connect.AddItem(), connect.EditItem(c.ServerIpAddr, c.LocalPort), connect.DelItem())
	widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(c), e.AbsolutePosition)
}

func (c *SelectClickText) DoubleTapped(e *fyne.PointEvent) {
	fmt.Println("double click at", e)
}
