package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// CLI is Command Line Interface.
type CLI struct {
	outStream io.Writer
}

// Run executes the CLI program.
func (c *CLI) Run(args []string) {
	fmt.Fprintln(c.outStream, strings.Join(args, " "))
}

func main() {
	cli := &CLI{outStream: os.Stdout}
	cli.Run(os.Args)
}
