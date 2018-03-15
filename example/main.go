package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gen0cide/bunnyhop"
)

var (
	// durFlag is used to set a timeout for an ARP request
	durFlag = flag.Duration("d", 1*time.Second, "timeout for ARP request")

	// ifaceFlag is used to set a network interface for ARP requests
	ifaceFlag = flag.String("i", "eth0", "network interface to use for ARP request")

	// ipFlag is used to set an IPv4 address destination for an ARP request
	ipFlag = flag.String("ip", "", "IPv4 address destination for ARP request")
)

func main() {
	flag.Parse()

	inUse, err := bunnyhop.IPInUse(*ifaceFlag, *ipFlag, *durFlag)

	if err != nil {
		log.Fatal(err)
	}

	if inUse == true {
		fmt.Printf("IN USE\n")
		return
	}

	fmt.Printf("AVAILABLE\n")
}
