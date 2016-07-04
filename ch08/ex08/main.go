package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	inputs := make(chan string)
	abort := make(chan struct{})

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputs <- input.Text()
		}
		abort <- struct{}{}
	}()

loop:
	for {
		select {
		case text := <-inputs:
			wg.Add(1)
			go func() {
				echo(c, text, 1*time.Second)
				wg.Done()
			}()
		case <-time.After(10 * time.Second):
			break loop
		case <-abort:
			break loop
		}
	}

	wg.Wait()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
