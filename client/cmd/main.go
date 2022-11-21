package main

import (
	"flag"
	"github.com/dreamlu/w2socks/client/core"
	"log"
	"os"
)

func main() {
	var i, p string
	flag.StringVar(&i, "i", "", "远程连接地址(包含端口)")
	flag.StringVar(&p, "p", "8018", "本地监听端口")
	flag.Parse()
	// 尝试从环境变量获取参数
	if i == "" {
		i = os.Getenv("IP_ADDR")
		if i == "" {
			log.Print("缺少远程地址")
			return
		}
	}
	log.Printf("local listen port: %s", p)
	log.Printf("server addr: %s", i)
	core.Core(&core.W2Config{
		ServerIpAddr: i,
		LocalPort:    p,
	})
}
