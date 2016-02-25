package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Exit codes
const (
	ExitCodeOK = iota
	ExitCodeHTTPError
	ExitCodeCopyError
)

// CLI is Command Line Interface.
type CLI struct {
	outStream, errStream io.Writer
}

// Run executes the echo program.
func (c *CLI) Run(args []string) int {
	for _, url := range args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(c.errStream, "fetch: %v\n", err)
			return ExitCodeHTTPError
		}
		defer resp.Body.Close()
		if _, err = io.Copy(c.outStream, resp.Body); err != nil {
			fmt.Fprintf(c.errStream, "fetch: reading %s: %v\n", url, err)
			return ExitCodeCopyError
		}
	}
	return ExitCodeOK
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
