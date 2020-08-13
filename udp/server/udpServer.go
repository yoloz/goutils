package main

import (
	"fmt"
	"net"
)

var addr string

func main() {
	fmt.Printf("please enter listen addr(ip:port): ")
	fmt.Scanf("%s", &addr)
	//创建监听的地址，并且指定udp协议
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr) //创建数据通信socket
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer conn.Close()
	for {
		buf := make([]byte, 4096)
		n, raddr, err := conn.ReadFromUDP(buf) //接收客户端发送过来的数据，填充到切片buf中。
		if err != nil {
			return
		}
		fmt.Printf("接受来自[%s]的数据[%s]\n", raddr.String(), string(buf[:n]))

		_, err = conn.WriteToUDP([]byte("nice to see u in udp"), raddr) // 向客户端发送数据
		if err != nil {
			fmt.Println("WriteToUDP err:", err)
			return
		}
	}

}
