package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tsamba120/go-socket-programming/pkg/helpers"
)

const (
	SERVER_TYPE = "tcp"
)

// function to be ran in a separate go routine to concurrently
// read any server responses in a new thread
func readServerResponse(conn net.Conn) {
	connScanner := bufio.NewScanner(conn)

	for connScanner.Scan() {
		fmt.Printf("%s\n", connScanner.Text())

		// if we encounter an error end the process (which, again, is a goroutine)
		if err := connScanner.Err(); err != nil {
			log.Fatalf(
				"Error reading server response from %s: %v",
				conn.LocalAddr().String(),
				err,
			)
		}
	}
}

// function to scan from StdIn and send it to the server
func sendStdInToServer(conn net.Conn) {
	for stdInScanner := bufio.NewScanner(os.Stdin); stdInScanner.Scan(); {
		// create byte array from stdin value and append newline char to end
		toBuild := [][]byte{}
		toBuild = append(toBuild, stdInScanner.Bytes())
		toBuild = append(toBuild, []byte("\n"))
		msg := bytes.Join(toBuild, nil)

		// send bytes from std in to server
		if _, err := conn.Write(msg); err != nil {
			log.Fatalf("Error writing to %s: %v", conn.LocalAddr().String(), err)
		}

		if err := stdInScanner.Err(); err != nil {
			log.Fatalf("Error reading from %s: %v", conn.LocalAddr().String(), err)
		}

	}
}

func main() {
	log.SetPrefix("tsamba_client:\t")

	addr := helpers.GetConnectionAddr()

	// establish connection
	conn, err := net.Dial(SERVER_TYPE, addr)
	if err != nil {
		log.Fatalf("error connecting to %s : %v", addr, err)
	}
	log.Printf(
		"Successfully created client. Connected to server at %s\n",
		conn.LocalAddr().String(),
	)
	defer conn.Close()

	// read server response in a separate thread
	go readServerResponse(conn)

	// read incoming lines from stdin and send them to server
	sendStdInToServer(conn)
}
