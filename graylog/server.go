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
	buf := make([]byte, 8192)
	for {
		n, err := sConn.Read(buf)

		data := check(buf[0:n])
		if data != "" {
			fmt.Println(data)
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
	//	defer sConn.Close()
}
