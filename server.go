package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	//=============
	ln, err := net.Listen("tcp", ":7788")
	checkError(err)

	for {
		conn, err := ln.Accept()
		checkError(err)
		go handleConn(conn)
	}
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("New connection!")
	msg, err := bufio.NewReader(conn).ReadString('\n')
	checkError(err)
	conn.Write([]byte(msg))
	conn.Close()
}
