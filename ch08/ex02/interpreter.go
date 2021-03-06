package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type interpreter struct {
	c    net.Conn
	cwd  string
	addr string
}

func newInterpreter(c net.Conn) *interpreter {
	return &interpreter{
		c:    c,
		cwd:  ".",
		addr: c.RemoteAddr().String(),
	}
}

func (ip *interpreter) reply(code string, text string) {
	io.WriteString(ip.c, fmt.Sprintf("%s %s\r\n", code, text))
}

func (ip *interpreter) start() {
	defer ip.c.Close()
	ip.reply("220", "Service ready for new user.")

	scanner := bufio.NewScanner(ip.c)
	for scanner.Scan() {
		command := scanner.Text()
		log.Println(command)

		if !ip.handleCommand(command) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("error:", err)
	}
}

var reUserCommand = regexp.MustCompile("(?i)^USER (.+?)$")
var reQuitCommand = regexp.MustCompile("(?i)^QUIT$")
var reTypeCommand = regexp.MustCompile("(?i)^TYPE (.+?)$")
var rePortCommand = regexp.MustCompile("(?i)^PORT (\\d+),(\\d+),(\\d+),(\\d+),(\\d+),(\\d+)$")
var reStruCommand = regexp.MustCompile("(?i)^STRU (.+?)$")
var reStorCommand = regexp.MustCompile("(?i)^STOR (.+?)$")
var reRetrCommand = regexp.MustCompile("(?i)^RETR (.+?)$")
var reNoopCommand = regexp.MustCompile("(?i)^NOOP$")
var reListCommand = regexp.MustCompile("(?i)^LIST(?: (.+?))?$")
var reCwdCommand = regexp.MustCompile("(?i)^CWD (.+?)$")

func (ip *interpreter) handleCommand(c string) (cont bool) {
	switch {
	case reUserCommand.MatchString(c):
		ip.handleUserCommand(c)
	case reTypeCommand.MatchString(c):
		ip.handleTypeCommand(c)
	case rePortCommand.MatchString(c):
		ip.handlePortCommand(c)
	case reStruCommand.MatchString(c):
		ip.handleStruCommand(c)
	case reRetrCommand.MatchString(c):
		ip.handleRetrCommand(c)
	case reStorCommand.MatchString(c):
		ip.handleStorCommand(c)
	case reListCommand.MatchString(c):
		ip.handleListCommand(c)
	case reCwdCommand.MatchString(c):
		ip.handleCwdCommand(c)
	case reNoopCommand.MatchString(c):
		ip.handleNoopCommand(c)
	case reQuitCommand.MatchString(c):
		ip.handleQuitCommand(c)
		return false
	default:
		ip.reply("502", "Command not implemented.")
	}
	return true
}

func (ip *interpreter) handleUserCommand(c string) {
	ip.reply("230", "User logged in, proceed.")
}

func (ip *interpreter) handleTypeCommand(c string) {
	ms := reTypeCommand.FindStringSubmatch(c)

	typeCode := ms[1]
	if typeCode != "A" && typeCode != "I" {
		ip.reply("504", "Command not implemented for that parameter.")
		return
	}
	ip.reply("200", "Command okay.")
}

func (ip *interpreter) handlePortCommand(c string) {
	ms := rePortCommand.FindStringSubmatch(c)
	p1, _ := strconv.Atoi(ms[5])
	p2, _ := strconv.Atoi(ms[6])

	host := fmt.Sprintf("%s.%s.%s.%s", ms[1], ms[2], ms[3], ms[4])
	port := p1*256 + p2
	ip.addr = fmt.Sprintf("%s:%d", host, port)
	ip.reply("200", "PORT command successful.")
}

func (ip *interpreter) handleStruCommand(c string) {
	// ms := reStruCommand.FindStringSubmatch(c)
	ip.reply("200", "STRU command successful.")
}

func (ip *interpreter) handleRetrCommand(c string) {
	ms := reRetrCommand.FindStringSubmatch(c)
	p := path.Join(ip.cwd, ms[1])

	if _, err := os.Stat(p); err != nil {
		ip.reply("550", "File not found.")
		return
	}

	ip.reply("150", "Open data connection.")
	conn, err := net.Dial("tcp", ip.addr)
	if err != nil {
		log.Println("error: ", err)
		ip.reply("425", "Can't open data connection.")
		return
	}
	defer conn.Close()

	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Println("error: ", err)
		ip.reply("550", "Can't open file.")
		return
	}

	if _, err = conn.Write(b); err != nil {
		log.Println("error: ", err)
		ip.reply("426", "Connection closed; transfer aborted.")
		return
	}

	ip.reply("250", "File transfer completed.")
}

func (ip *interpreter) handleStorCommand(c string) {
	ms := reStorCommand.FindStringSubmatch(c)
	p := path.Join(ip.cwd, ms[1])

	ip.reply("150", "File status okay; about to open data connection.")

	conn, err := net.Dial("tcp", ip.addr)
	if err != nil {
		log.Println("error: ", err)
		ip.reply("425", "Can't open data connection.")
		return
	}
	defer conn.Close()

	f, err := os.Create(p)
	if err != nil {
		log.Println("error: ", err)
		ip.reply("450", "Requested file action not taken.")
		return
	}
	defer f.Close()

	if _, err = io.Copy(f, conn); err != nil {
		log.Println("error: ", err)
		ip.reply("450", "Requested file action not taken.")
		return
	}

	ip.reply("250", "File transfer completed.")
}

func (ip *interpreter) handleNoopCommand(c string) {
	ip.reply("200", "NOOP command successful.")
}

func (ip *interpreter) handleQuitCommand(c string) {
	ip.reply("221", "Service closing control connection.")
}

func (ip *interpreter) handleCwdCommand(c string) {
	ms := reCwdCommand.FindStringSubmatch(c)
	p := ms[1]

	if !strings.HasPrefix(p, "/") {
		p = path.Join(ip.cwd, p)
	}

	info, err := os.Stat(p)
	if err != nil || !info.IsDir() {
		ip.reply("550", "Directory not found.")
		return
	}
	ip.cwd = p

	ip.reply("250", "Requested file action okay, completed.")
}

func (ip *interpreter) handleListCommand(c string) {
	ms := reListCommand.FindStringSubmatch(c)
	p := path.Join(ip.cwd, ms[1])

	ip.reply("150", "Open data connection.")
	conn, err := net.Dial("tcp", ip.addr)
	if err != nil {
		log.Println("error: ", err)
		ip.reply("425", "Can't open data connection.")
		return
	}
	defer conn.Close()

	info, err := os.Stat(p)
	if err != nil {
		ip.reply("550", "File or Directory not found.")
		return
	}

	var result []byte
	if info.IsDir() {
		is, err := ioutil.ReadDir(p)
		if err != nil {
			ip.reply("550", "File or Directory not found.")
			return
		}

		for _, i := range is {
			result = append(result, []byte(i.Name()+"\r\n")...)
		}
	} else {
		result = []byte(fmt.Sprintf("%s\r\n", p))
	}

	if _, err = conn.Write(result); err != nil {
		log.Println("error: ", err)
		ip.reply("426", "Connection closed; transfer aborted.")
		return
	}

	ip.reply("250", "File transfer completed.")
}
