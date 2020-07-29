package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/util/ip"
	"github.com/dreamlu/w2socks/client/util/notify"
	"github.com/kardianos/service"
	"log"
)

var logger service.Logger

// 运行方式:
// 1.命令行
// 2.GUI
func main() {

	window()

	svcConfig := &service.Config{
		Name:        "w2socks-program",
		DisplayName: "w2socks",
		Description: "This is a proxy Go service.",
	}

	prg := &BackstageProgram{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		err = logger.Error(err)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

// window
func window() {
	// 主程序
	majorApp := app.New()

	// logo
	majorApp.SetIcon(data.Logo())
	// 主窗体
	mainWindow := majorApp.NewWindow("w2socks")
	mainWindow.Resize(fyne.NewSize(280, 300))

	comSize := fyne.NewSize(100, 20)

	// 服务端ip和端口
	serverEntry := widget.NewEntry()
	serverEntry.SetPlaceHolder("ip:port")
	serverEntry.Resize(comSize)

	// 本地端口号
	localPortEntry := widget.NewEntry()
	localPortEntry.SetPlaceHolder("port")
	localPortEntry.Resize(comSize)

	form := widget.NewForm(
		widget.NewFormItem("server:", serverEntry),
		widget.NewFormItem("local:", localPortEntry),
	)

	form.CancelText = "disconnect"
	form.SubmitText = "connect"
	// 取消操作
	form.OnCancel = func() {
		// 退出旧携程
		if core.Online > 0 {
			core.Quit <- 1
			log.Println("取消")
			notify.SysNotify("notify", "server is disconnected")
		}
	}
	// 连接操作
	form.OnSubmit = func() {
		log.Println("提交")
		ipAddr := serverEntry.Text
		log.Println("用户输入: " + ipAddr)

		// ip地址是否正确
		msg, ok := ip.Check(ipAddr)
		if !ok {
			notify.SysNotify("warn!!", msg)
			return
		}

		//本地端口是否正确
		if !ip.CheckPort(localPortEntry.Text) {
			notify.SysNotify("warn!!", "Incorrect local port")
			return
		}

		// 退出旧携程
		if core.Online > 0 {
			core.Quit <- 1
		}
		go core.Core(ipAddr, localPortEntry.Text)
		core.Online++
		// 系统通知
		notify.SysNotify("w2socks", "success to connect "+ipAddr)
	}

	// 窗体
	content := widget.NewVBox(
		widget.NewVBox(
			// 输入服务端的ip地址和端口 以及本地的端口
			widget.NewLabel("Please Enter:"),
			form,
		),
	)
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

// 后台
type BackstageProgram struct {
}

func (p *BackstageProgram) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *BackstageProgram) run() {
	// Do work here
}

func (p *BackstageProgram) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
