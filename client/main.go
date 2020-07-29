package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dreamlu/w2socks/client/core"
	"github.com/dreamlu/w2socks/client/data"
	"github.com/dreamlu/w2socks/client/util/notify"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// 运行方式:
// 1.命令行
// 2.GUI
func main() {
	window()
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
		if !Check(ipAddr) {
			return
		}

		//本地端口是否正确
		if !CheckPort(localPortEntry.Text) {
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

func Check(ipAddr string) bool {
	if !strings.Contains(ipAddr, ":") {
		fmt.Println("ip和端口格式不正确")
		notify.SysNotify("warn!!", "ip and port format is incorrect")
		return false
	}
	ip := strings.Split(ipAddr, ":")
	ipv4 := ip[0]
	if !CheckIp(ipv4) {
		fmt.Println("ip地址格式不正确")
		notify.SysNotify("warn!!", "ip地址格式不正确")
		return false
	}
	port := ip[1]
	if !CheckPort(port) {
		fmt.Println("")
		notify.SysNotify("warn!!", "ip端口不正确")
		return false
	}
	return true
}

// 检验ip地址
func CheckIp(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func CheckPort(port string) bool {
	portNum, err := strconv.Atoi(port)
	if err != nil || portNum >= 65536 || portNum <= 0 {
		return false
	}
	return true
}
