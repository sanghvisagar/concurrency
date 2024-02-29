package main

import (
	"fmt"
	"log"
	"net"
)

func processConn(conn net.Conn) {
	var request []byte

	conn.Read(request)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))

	conn.Close()

}

func main() {

	listener, err := net.Listen("tcp", ":1729")

	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Conn accepted")

		go processConn(conn)
	}

}
