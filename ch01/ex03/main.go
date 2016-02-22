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

// Run executes the echo program.
func (c *CLI) Run(args []string) {
	fmt.Fprintln(c.outStream, strings.Join(args[1:], " "))
}

// RunInefficiently executes the echo program inefficiently.
func (c *CLI) RunInefficiently(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(c.outStream, s)
}

func main() {
	cli := &CLI{outStream: os.Stdout}
	cli.Run(os.Args)
}
