package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Printf("Running server on")
	server, err := net.Listen("tcp", SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Printf("Listening on %s:%s\n", SERVER_HOST, SERVER_PORT)
	fmt.Println("Waiting on client...")

	for {
		clientConnection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Client accepted")
		// handle connection with go routine
		go processClientMessage(clientConnection)
	}
}

func processClientMessage(connection net.Conn) {
	buffer := make([]byte, 1024)
	bytesRead, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read from connection:", err.Error())
	}

	msg := string(buffer[:bytesRead])
	fmt.Println("Received: ", msg)

	_, err = connection.Write([]byte("Thanks! Got your message: " + msg))
	if err != nil {
		fmt.Println("Failed to write back to client: ", err.Error())
	}
}
