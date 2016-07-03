package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port int

func init() {
	const (
		portDefault = 8000
		portUsage   = "port"
	)

	flag.IntVar(&port, "port", portDefault, portUsage)
	flag.IntVar(&port, "p", portDefault, portUsage+" (shorthand)")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		ip := newInterpreter(conn)
		go ip.start()
	}
}
