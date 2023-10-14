package helpers

import (
	"flag"
)

// function to grab CLI arguments for which server to connect to
func GetConnectionAddr() string {
	server_host_ptr := flag.String("host", "localhost", "server host name")
	server_port_ptr := flag.String("port", "8080", "server port to connect to")
	return *server_host_ptr + ":" + *server_port_ptr
}