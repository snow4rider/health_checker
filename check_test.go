package main

import (
	"net"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	t.Run("CheckDestinationPort", func(t *testing.T) {
		destination := "www.google.com"
		port := "80"
		timeout := 5 * time.Second

		conn, err := net.DialTimeout("tcp", destination+":"+port, timeout)
		if err != nil {
			t.Fatalf("[DOWN] %s is unreachable, Error: %v", destination, err)
		} else {
			t.Logf("[UP] %s is reachable, From: %v To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
		}
	})

}
