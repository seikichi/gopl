package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
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
	c    net.Conn
	cwd  string
	host string
	port int
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

var reUserCommand = regexp.MustCompile("(?i)^USER (.*?)$")
var reTypeCommand = regexp.MustCompile("(?i)^TYPE (.*?)$")
var rePortCommand = regexp.MustCompile("(?i)^PORT (\\d+),(\\d+),(\\d+),(\\d+),(\\d+),(\\d+)$")
var reRetrCommand = regexp.MustCompile("(?i)^RETR (.*?)$")
var reNoopCommand = regexp.MustCompile("(?i)^NOOP$")

func (ip *interpretor) handleCommand(c string) {
	switch {
	case reUserCommand.MatchString(c):
		ip.handleUserCommand(c)
	case reTypeCommand.MatchString(c):
		ip.handleTypeCommand(c)
	case rePortCommand.MatchString(c):
		ip.handlePortCommand(c)
	case reRetrCommand.MatchString(c):
		ip.handleRetrCommand(c)
	case reNoopCommand.MatchString(c):
		ip.handleNoopCommand(c)
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

func (ip *interpretor) handlePortCommand(c string) {
	ms := rePortCommand.FindStringSubmatch(c)
	p1, _ := strconv.Atoi(ms[5])
	p2, _ := strconv.Atoi(ms[6])

	ip.host = fmt.Sprintf("%s.%s.%s.%s", ms[1], ms[2], ms[3], ms[4])
	ip.port = p1*256 + p2
	io.WriteString(ip.c, "200 PORT command successful.\r\n")
}

func (ip *interpretor) handleRetrCommand(c string) {
	ms := reRetrCommand.FindStringSubmatch(c)
	p := path.Join(ip.cwd, ms[1])

	if _, err := os.Stat(p); err != nil {
		io.WriteString(ip.c, "550 File not found.\r\n")
		return
	}

	io.WriteString(ip.c, "150 File status okay; about to open data connection.\r\n")
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip.host, ip.port))
	if err != nil {
		log.Println("error: ", err)
		io.WriteString(ip.c, "425 Can't open data connection.\r\n")
		return
	}
	defer conn.Close()

	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Println("error: ", err)
		io.WriteString(ip.c, "550 Can't open file.\r\n")
		return
	}

	_, err = conn.Write(b)
	if err != nil {
		log.Println("error: ", err)
		io.WriteString(ip.c, "426 Connection closed; transfer aborted.\r\n")
		return
	}

	io.WriteString(ip.c, "250 File transfer completed.\r\n")
}

func (ip *interpretor) handleNoopCommand(c string) {
	io.WriteString(ip.c, "200 NOOP command successful.\r\n")
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
		ip := &interpretor{c: conn, cwd: "."}
		go ip.start()
	}
}
