package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:12201")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("YourDataHere")); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	//	buf := []byte(b)
	_, err = Conn.Write(b.Bytes())

	//    fmt.Println(b)
	//	i := 0
	//	for {
	//		msg := strconv.Itoa(i)
	//		i++
	//		buf := []byte(msg)
	//		_, err := Conn.Write(buf)
	//		if err != nil {
	//			fmt.Println(msg, err)
	//		}
	//		time.Sleep(time.Second * 1)
	//	}
}
