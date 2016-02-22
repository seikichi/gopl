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
	lines := []string{}
	for i, arg := range args[1:] {
		lines = append(lines, fmt.Sprintf("%d\t%s", i+1, arg))
	}
	fmt.Fprintln(c.outStream, strings.Join(lines, "\n"))
}

func main() {
	cli := &CLI{outStream: os.Stdout}
	cli.Run(os.Args)
}
