package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
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

type interpretor struct {
	c net.Conn
}

func (ip *interpretor) start() {
	defer ip.c.Close()
	io.WriteString(ip.c, "220\r\n")

	scanner := bufio.NewScanner(ip.c)
	for scanner.Scan() {
		command := scanner.Text()
		log.Println(command)
		ip.handleCommand(command)
	}
	if err := scanner.Err(); err != nil {
		log.Println("error:", err)
	}
}

var reUserCommand = regexp.MustCompile("^USER (.*?)$")
var reTypeCommand = regexp.MustCompile("^TYPE (.*?)$")

func (ip *interpretor) handleCommand(c string) {
	switch {
	case reUserCommand.MatchString(c):
		ip.handleUserCommand(c)
	case reTypeCommand.MatchString(c):
		ip.handleTypeCommand(c)
	default:
		io.WriteString(ip.c, "502 Command not implemented.\r\n")
	}
}

func (ip *interpretor) handleUserCommand(c string) {
	io.WriteString(ip.c, "230 User logged in, proceed.\r\n")
}

func (ip *interpretor) handleTypeCommand(c string) {
	ms := reTypeCommand.FindStringSubmatch(c)

	typeCode := ms[1]
	if typeCode != "A" && typeCode != "I" {
		io.WriteString(ip.c, "504 Command not implemented for that parameter.\r\n")
		return
	}
	io.WriteString(ip.c, "200\r\n")
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
		ip := &interpretor{c: conn}
		go ip.start()
	}
}
