package main

import (
	"fmt"
	"net"
)

func main() {
	serverConnection, err := net.Dial("tcp", "localhost:9999")

	if err != nil {
		return
	}
	defer serverConnection.Close()

	payload := "Hello World!"
	_, err = serverConnection.Write([]byte(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
}
