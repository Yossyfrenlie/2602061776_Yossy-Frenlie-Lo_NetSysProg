package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		return
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			return
		}
		go func(client net.Conn) {
			fmt.Println("Accepted")
			buf := make([]byte, 100)
			n, err := client.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("read: %d, message %s", n, buf)
		}(clientConn)

	}

}
