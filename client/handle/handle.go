package handle

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
)

// 本机请求处理
// 处理本机的数据并发送请求
func Send(client net.Conn, c *websocket.Conn) error {

	//defer client.Close()
	b := make([]byte, 1024)
	for {
		readCount, errRead := client.Read(b)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			// 发送数据
			writeCount, err := c.UnderlyingConn().Write(b[0:readCount]) //c.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				//client.Close()
				log.Println("write err:", err)
				return err
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

func Receive(client net.Conn, c *websocket.Conn) error {

	//defer client.Close()
	b := make([]byte, 1024)
	for {
		readCount, errRead := c.UnderlyingConn().Read(b)
		if errRead != nil {
			return errRead
		}
		if readCount > 0 {
			// 写回本地
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
