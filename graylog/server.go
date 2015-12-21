package main

import (
	"fmt"
	"net"
)

func main() {
	sAddr, err := net.ResolveUDPAddr("udp", ":12201")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	sConn, err := net.ListenUDP("udp", sAddr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//	fmt.Println("listening on ", sConn.LocalAddr().String())
	buf := make([]byte, 1024)
	for {
		n, err := sConn.Read(buf)

		fmt.Println(check(buf[0:n]))

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
	//	defer sConn.Close()
}
