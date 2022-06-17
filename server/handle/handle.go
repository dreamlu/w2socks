package handle

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
)

// Handle socks5: https://jiajunhuang.com/articles/2019_06_06-socks5.md.html
func Handle(ws *websocket.Conn) {

	//defer ws.Close()
	client := ws.UnderlyingConn()
	b := make([]byte, 256)
	// 读取本机的代理请求内容
	_, err := client.Read(b)
	if err != nil {
		log.Println(err)
		return
	}
	if b[0] != 0x05 {
		log.Println("error: not socks5 protocol")
		return
	}

	// 告诉client无需认证
	_, err = client.Write([]byte{0x05, 0x00})
	if err != nil {
		log.Println(fmt.Sprintf("error:%s", err))
		return
	}

	// 获取真正的远程服务的地址
	n, err := client.Read(b)
	// n 最短的长度为7 情况为 ATYP=3 ipv4值为0x01, DST.ADDR占用4字节
	if err != nil || n < 7 {
		return
	}
	var addr []byte
	// ATYP 是目标地址类型，有如下取值：
	// 0x01 IPv4
	// 0x03 域名
	// 0x04 IPv6
	// DST.ADDR 就是目标地址的值了，如果是IPv4，那么就是4 bytes，如果是IPv6那么就是16 bytes，如果是域名，那么第一个字节代表 接下来有多少个字节是表示目标地址
	switch b[3] {
	case 0x01:
		//	IP V4 address: X'01'
		addr = b[4 : 4+net.IPv4len]
	case 0x03:
		//	DOMAINNAME: X'03'
		ipAddr, err := net.ResolveIPAddr("ip", string(b[5:n-2]))
		if err != nil {
			return
		}
		addr = ipAddr.IP
	case 0x04:
		//	IP V6 address: X'04'
		addr = b[4 : 4+net.IPv6len]
	default:
		return
	}

	dPort := b[n-2:]
	dstAddr := &net.TCPAddr{
		IP:   addr,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}

	// CMD代表客户端请求的类型，值长度也是1个字节，有三种类型
	// CONNECT X'01'
	if b[1] != 0x01 {
		// 目前只支持 CONNECT
		return
	}

	// ===== 建立连接 ====
	// 连接服务
	server, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		_ = server.SetLinger(0)
		server.Close()
		client.Close()
	}()
	// Conn被关闭时直接清除所有数据 不管没有发送的数据
	// 响应客户端连接成功
	_, _ = client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	// 进行转发
	// 从server接收数据,返回本地
	go io.Copy(server, client)
	_, err = io.Copy(client, server)
	if err == nil {
		client.Write([]byte(io.EOF.Error()))
	}
}

// 源源不断的接收数据
//func receive(client io.WriteCloser, server *net.TCPConn) error {
//	b := make([]byte, 1024)
//	for {
//		readCount, errRead := server.Read(b)
//		if errRead != nil {
//			if errRead != io.EOF {
//				return errRead
//			}
//			return nil
//		}
//		if readCount > 0 {
//			writeCount, errWrite := client.Write(b[0:readCount])
//			if errWrite != nil {
//				return errWrite
//			}
//			if readCount != writeCount {
//				return io.ErrShortWrite
//			}
//		}
//	}
//}
//
//// 源源不断的发送数据
//func send(client io.Reader, server *net.TCPConn) error {
//
//	b := make([]byte, 1024)
//	for {
//		readCount, errRead := client.Read(b)
//		if errRead != nil {
//			if errRead != io.EOF {
//				return errRead
//			}
//			return nil
//		}
//		if readCount > 0 {
//			writeCount, errWrite := server.Write(b[0:readCount])
//			if errWrite != nil {
//				return errWrite
//			}
//			if readCount != writeCount {
//				return io.ErrShortWrite
//			}
//		}
//	}
//}
