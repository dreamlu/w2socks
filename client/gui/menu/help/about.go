package help

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"log"
	"net/url"
)

// 添加逻辑
func HelpItem() *fyne.MenuItem {
	return fyne.NewMenuItem("About", func() {
		w := fyne.CurrentApp().NewWindow("About")
		w.SetContent(widget.NewHyperlink("w2socks", parseURL("https://github.com/dreamlu/w2socks")))
		w.Show()
	})
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Could not parse URL", err)
	}

	return link
}
