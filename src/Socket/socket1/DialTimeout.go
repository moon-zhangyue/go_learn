package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	// 从参数中读取主机信息
	service := os.Args[1]

	// 建立网络连接 3秒超时
	conn, err := net.DialTimeout("tcp", service, 3*time.Second)
	// 连接出错则打印错误消息并退出程序
	checkError(err)

	// 调用返回的连接对象提供的 Write 方法发送请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 通过连接对象提供的 Read 方法读取所有响应数据
	result, err := readFully(conn)
	checkError(err)

	// 设置读写超时时间
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	checkError(err)

	// 打印响应数据
	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	// 读取所有响应数据后主动关闭连接
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
