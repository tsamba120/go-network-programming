package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func main() {
	// establish connection
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello world"))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)

	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	} else {
		fmt.Printf("Read %d bytes\n", bytesRead)
		fmt.Printf("Message: %s\n", string(buffer[:bytesRead]))
	}

}
