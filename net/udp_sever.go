package main

import (
	"fmt"
	"net"
)

func main() {
	//listen on port 7070
	addr, err := net.ResolveUDPAddr("udp", ":7070")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleClient(conn, buf[0:n], addr)
	}

}

func HandleClient(conn *net.UDPConn, data []byte, addr *net.UDPAddr) {
	fmt.Println("收到数据:" + string(data))
	conn.WriteToUDP([]byte("OK,数据收到"), addr)
}
