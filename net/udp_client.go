package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7070")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello Server"))
	var buf [1024]byte
	n, _, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[0:n]))

}
