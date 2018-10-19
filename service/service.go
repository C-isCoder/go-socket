package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// TITLE is replay ack title.
const TITLE = "server-ack: "

func main() {
	service := ":6666"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 128)
	for {
		readLine, err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}
		if readLine == 0 {
			break
		} else {
			fmt.Printf("<<<------%s", b)
			fmt.Printf("<<<------%d", b[:readLine])
		}
	}
	daytime := time.Now().String()
	conn.Write([]byte(TITLE + daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
