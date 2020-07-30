package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"image/color"
)

type SelectClickText struct {
	canvas.Text
}

func NewSelectClick(content string, c color.Color) *SelectClickText {
	sel := SelectClickText{}
	sel.Text.Text = content
	sel.Color = c
	return &sel
}

func (c *SelectClickText) Tapped(e *fyne.PointEvent) {
	fmt.Println("left click at", e)
}

func (c *SelectClickText) TappedSecondary(e *fyne.PointEvent) {
	fmt.Println("right click at", e)
}

func (c *SelectClickText) DoubleTapped(e *fyne.PointEvent) {
	fmt.Println("double click at", e)
}
