package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// ACK is replay ack title.
const ACK = "server-ack:(OK)"

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
			fmt.Printf("<<<------%s-time:%s", b, time.Now().Format("2006-01-02 15:04:05"))
			conn.Write([]byte(ACK))
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
