package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// flags
var (
	host string
	path string
	method string
	port int
)

func main() {

	flag.StringVar(&method, "method", "GET", "HTTP method to use")
	flag.StringVar(&host, "host", "localhost", "Host to connect to")
	flag.IntVar(&port, "port", 80, "Port to connect to")
	flag.StringVar(&path, "path", "/", "Path to request")
	flag.Parse()

	// ResolveTCPAddr is slightly more convenient way of creating a TCPAddr.
	// Let's use this since we've already done net.LookupIP by hand
	ip, err := net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf("%s:%d", host, port),
	)
	if err != nil {
		panic(err)
	}

	// dial the remove host using the TCPAddr we just created....
	conn, err := net.DialTCP("tcp", nil, ip)
	if err != nil {
		panic(err)
	}

	log.Printf("connected to %s (@ %s)", host, conn.RemoteAddr())

	defer conn.Close()

	// Build request string
	// e.g, for a request to http://eblog.fly.dev/
    // GET / HTTP/1.1
    // Host: eblog.fly.dev
    // User-Agent: httpget
    //
	var reqfields = []string{
		fmt.Sprintf("%s %s HTTP/1.1", method, path),
		"Host: " + host,
		"User-Agent: httpget",
		"", // Empty line to terminate the headers
	}

	request := strings.Join(reqfields, "\n") + "\n"

	// send request
	conn.Write([]byte(request))
	log.Printf("sent request: \n%s\n", request)

	for scanner := bufio.NewScanner(conn); scanner.Scan(); {
		line := scanner.Bytes()
		if _, err := fmt.Fprintf(os.Stdout, "%s\n", line); err != nil {
			log.Printf("error writing to connection: %s", err)
		}
		if err = scanner.Err(); err != nil {
			log.Printf("error reading from connection: %s", err)
			return
		}
	}
}