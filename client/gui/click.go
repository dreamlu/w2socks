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

	addItem := connect.AddItem()
	editItem := connect.EditItem()
	delItem := connect.DelItem()

	var menu = fyne.NewMenu("", addItem, editItem, delItem)

	entryPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(c)
	popUpPos := entryPos.Add(fyne.NewPos(e.Position.X, e.Position.Y))
	d := fyne.CurrentApp().Driver().CanvasForObject(c)

	widget.ShowPopUpMenuAtPosition(menu, d, popUpPos)
}

func (c *SelectClickText) DoubleTapped(e *fyne.PointEvent) {
	fmt.Println("double click at", e)
}
