package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/tsamba120/go-socket-programming/pkg/helpers"
)

const (
	SERVER_TYPE = "tcp"
)

// TODO: learn how to automatically format and lint files

// function to process a client message in a new thread
func processClientMessage(connection net.Conn) {
	for {
		buffer := make([]byte, 1024)
		bytesRead, err := connection.Read(buffer)
		if err != nil {
			log.Printf("ERROR: \t Failed to read from connection: %v\n", err.Error())
			break
		}

		msg := string(buffer[:bytesRead])
		fmt.Println("Received from client: ", msg)

		serverResponse := fmt.Sprintf("Message received: %s\n", strings.ToUpper(msg))
		_, err = connection.Write([]byte(serverResponse))
		if err != nil {
			log.Println("ERROR: \t Failed to write back to client: ", err.Error())
		}
	}
}

func main() {
	log.SetPrefix("tsamba_server:\t")
	addr := helpers.GetConnectionAddr()

	fmt.Printf("Running server on")
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listening: %v", err.Error())
	}
	defer server.Close()

	log.Printf("Listening on %s\n", addr)
	log.Println("Waiting on client...")

	// infinite loop, accepting connections from N number of clients
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
