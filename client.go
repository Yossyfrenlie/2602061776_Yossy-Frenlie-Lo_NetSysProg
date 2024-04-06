package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert Message: ")
	scan.Scan()
	message := scan.Text()

	sendToServer(message)
}

func sendToServer(message string) {
	serverConnection, err := net.DialTimeout("tcp", "127.0.0.1:9999", 15*time.Second)
	if err != nil {
		panic(err)
	}
	defer serverConnection.Close()

	err = binary.Write(serverConnection, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}

	_, err = serverConnection.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	var size uint32
	err = binary.Read(serverConnection, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}

	byteReply := make([]byte, size)
	serverConnection.SetReadDeadline(time.Now().Add(15 * time.Second))
	_, err = serverConnection.Read(byteReply)
	if err != nil {
		panic(err)
	}
	reply := string(byteReply)
	fmt.Printf("Replied: %s\n", reply)
}
