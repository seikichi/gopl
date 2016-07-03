package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

type item struct {
	key  string
	time string
}

func main() {
	c := make(chan *item)
	keys := []string{}
	times := map[string]string{}

	for _, arg := range os.Args[1:] {
		sp := strings.Split(arg, "=")
		if len(sp) != 2 {
			log.Fatalf("invalid argument: %q", arg)
		}

		key, addr := sp[0], sp[1]
		times[key] = ""
		keys = append(keys, key)

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}

		go func(key string, conn net.Conn, c chan<- *item) {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				c <- &item{key: key, time: scanner.Text()}
			}
		}(key, conn, c)
	}

	for i := range c {
		times[i.key] = i.time

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		for _, key := range keys {
			fmt.Printf("%s\t%s\n", key, times[key])
		}
	}
}
