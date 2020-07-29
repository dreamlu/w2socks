package ip

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Check(ipAddr string) (string, bool) {
	if !strings.Contains(ipAddr, ":") {
		fmt.Println("ip和端口格式不正确")
		return "Ip and port format is incorrect", false
	}
	ip := strings.Split(ipAddr, ":")
	ipv4 := ip[0]
	if !CheckIp(ipv4) {
		fmt.Println("ip地址格式不正确")
		return "Ip address format is incorrect", false
	}
	port := ip[1]
	if !CheckPort(port) {
		fmt.Println("")
		return "Ip port is incorrect", false
	}
	return "", true
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
