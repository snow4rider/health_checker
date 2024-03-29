package main

import (
	"fmt"
	"net"
	"time"
)

// CheckDestinationPort checks the reachability of a destination on a specific port.
//
// destination string - the destination address
// port string - the port number
// string - the status message indicating the reachability of the destination on the specified port
func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, From: %v To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
