package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"
	"strings"
)

type client chan<- string
type enter struct {
	client client
	name   string
}

var (
	entering = make(chan enter)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]string)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case enter := <-entering:
			names := []string{}
			for _, n := range clients {
				names = append(names, n)
			}
			sort.Strings(names)
			clients[enter.client] = enter.name
			enter.client <- "current clients: " + strings.Join(names, ", ")

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- enter{client: ch, name: who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
