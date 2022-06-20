package notify

import (
	"fyne.io/fyne/v2"
	"log"
)

// 全局通知方法

// 系统通知
func SysNotify(title, content string) {
	log.Println(content)
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   title,
		Content: content,
	})
}
