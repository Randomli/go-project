package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func RunClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		iniputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(iniputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err := conn.Write([]byte(iniputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n])) // 打印接收到的数据
	}
}
