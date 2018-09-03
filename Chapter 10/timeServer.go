package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// We could run on any port. Here we are running on 13 to match the DAYTIME
	// standard as described in RFC 867.
	port := ":13"
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Fatalf("resolve addr: %+v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen: %+v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connection open: %+v", err)
			continue
		}

		go Handle(conn)
	}
}

// Handle does two things, log when it receives a connection, and writes the
// current time to the connection and closes it.
func Handle(conn net.Conn) {
	now, err := time.Now().MarshalText()
	if err != nil {
		log.Printf("get time: %+v", err)
	}

	conn.Write(now)
	conn.Close()
}
