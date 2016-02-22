package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// CLI is Command Line Interface.
type CLI struct {
	outStream, errStream io.Writer
}

// Run executes the echo program.
func (c *CLI) Run(args []string) {
	for _, url := range args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(c.errStream, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if _, err = io.Copy(c.outStream, resp.Body); err != nil {
			fmt.Fprintf(c.errStream, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	cli.Run(os.Args)
}
