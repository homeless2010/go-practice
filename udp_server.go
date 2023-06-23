package main

import (
	"fmt"
	"net"
)

func main() {
	listenUDP, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	defer listenUDP.Close()
	for {
		var buf [1024]byte
		n, addr, err := listenUDP.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("red from udp failed,err:%v\n", err)
			return
		}
		fmt.Println("接收到的数据：", string(buf[:n]))
		_, err = listenUDP.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to %v failed,err:%v\n", addr, err)
			return
		}
	}
}
