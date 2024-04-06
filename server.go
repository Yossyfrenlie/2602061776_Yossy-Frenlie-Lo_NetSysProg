package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		clientConnection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleServer(clientConnection)
	}
}

func handleServer(client net.Conn) {
	var size uint32
	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytMsg := make([]byte, size)
	client.SetReadDeadline(time.Now().Add(15 * time.Second))
	_, err = client.Read(bytMsg)
	if err != nil {
		panic(err)
	}
	strMsg := string(bytMsg)
	fmt.Printf("Received: %s\n", strMsg)

	reply := "Message has been received"

	err = binary.Write(client, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}
	client.SetWriteDeadline(time.Now().Add(15 * time.Second))
	_, err = client.Write([]byte(reply))
	if err != nil {
		panic(err)
	}
}
