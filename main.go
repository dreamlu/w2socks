package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

func main() {

	// 启动监听
	LocalListenAddr, err := net.ResolveTCPAddr("tcp", ":8081")
	if err != nil {
		log.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", LocalListenAddr)
	if err != nil {
		log.Panic(err)
	}
	defer l.Close()

	for {
		// 监听对于本机的请求
		client, err := l.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("本机请求")
		// localConn被关闭时直接清除所有数据 不管没有发送的数据
		_ = client.SetLinger(0)
		// 处理本机的代理请求
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {

	defer client.Close()
	b := make([]byte, 256)

	// 获取真正的远程服务的地址
	n, err := client.Read(b)
	// n 最短的长度为7 情况为 ATYP=3 DST.ADDR占用1字节 值为0x0
	if err != nil || n < 7 {
		return
	}
	log.Println(b)

	var dIP []byte
	// aType 代表请求的远程服务器地址类型，值长度1个字节，有三种类型
	switch b[3] {
	case 0x01:
		//	IP V4 address: X'01'
		dIP = b[4 : 4+net.IPv4len]
	case 0x03:
		//	DOMAINNAME: X'03'
		ipAddr, err := net.ResolveIPAddr("ip", string(b[5:n-2]))
		if err != nil {
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		//	IP V6 address: X'04'
		dIP = b[4 : 4+net.IPv6len]
	default:
		return
	}

	dPort := b[n-2:]
	dstAddr := &net.TCPAddr{
		IP:   dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}

	// CMD代表客户端请求的类型，值长度也是1个字节，有三种类型
	// CONNECT X'01'
	if b[1] != 0x01 {
		// 目前只支持 CONNECT
		return
	}

	////获得了请求的host和port，就开始拨号吧
	//requestAddr, err := net.ResolveTCPAddr("tcp", req.Address)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	// ===== 进行连接 ====
	// 连接服务
	server, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	//if req.Method == "CONNECT" {
	//	fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n")
	//} else {
	//	server.Write(b[:n])
	//}
	// Conn被关闭时直接清除所有数据 不管没有发送的数据
	_ = server.SetLinger(0)
	// 响应客户端连接成功
	client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	// 响应客户端连接成功
	//server.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	// 进行转发
	// 从server接收数据,返回本地
	go send(client, server)
	// 从server->client,渲染
	receive(client, server)
	log.Println("转发结束")
}

// 源源不断的接收数据
func receive(client net.Conn, server *net.TCPConn) error {
	b := make([]byte, 1024)
	for {
		readCount, errRead := server.Read(b)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := client.Write(b[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// 源源不断的发送数据
func send(client net.Conn, server *net.TCPConn) error {

	b := make([]byte, 1024)
	for {
		readCount, errRead := client.Read(b)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := server.Write(b[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}
