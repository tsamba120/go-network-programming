package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// given host we look up it's IP address
func main() {
	if len(os.Args) != 2 {
		log.Printf("%s: usage: <host>", os.Args[0])
		log.Fatalf("expected exactly one argument; got %d", len(os.Args)-1)
	}
	host := os.Args[1]
	ipAddr, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("lookup ip: %s %s", host, err)
	} 
	if len(ipAddr) == 0 {
		log.Fatalf("no ip found for %s", host)
	}

	var found bool
	for _, ip := range ipAddr {
		if ip.To4() != nil {
			fmt.Println(ip)
			found = true
			break
		}
	}
	if found {
		checkIPV6(ipAddr)
	} else {
		fmt.Printf("none\n")
	}

}

func checkIPV6(ipAddr []net.IP) {
	for _, ip := range ipAddr {
		if ip.To4() == nil {
			fmt.Println(ip)
			return
		}
	}
	fmt.Printf("none\n")
}
