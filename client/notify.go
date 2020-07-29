package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/widget"
	"time"
)

// 全局通知方法

// 弹窗
func PopUps(app fyne.App, title, content string) {
	w := app.NewWindow(title)

	w.SetContent(widget.NewVBox(
		widget.NewLabel(content),
		widget.NewButton("Quit", func() {
			w.Close()
		}),
	))
	w.Show()
}

// 弹窗
func PopUpsWarn(app fyne.App, content string) {
	PopUps(app, "Warn!!", content)
}

// 弹窗
func PopSplashWarn(w fyne.Window, content string) {
	PopSplashNotify(w, "Warn!!", content)
}

// 系统通知
func SysNotify(title, content string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   title,
		Content: content,
	})
}

func PopSplashNotify(w fyne.Window, title, content string) {
	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		w.SetContent(
			widget.NewVBox(
				widget.NewButton(title, func() {
					w := drv.CreateSplashWindow()
					w.SetContent(widget.NewLabelWithStyle(content,
						fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
					w.Show()

					go func() {
						time.Sleep(time.Second * 3)
						w.Close()
					}()
				}),
			),
		)
	}
}
