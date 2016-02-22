package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// CLI is Command Line Interface.
type CLI struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

// Run executes the CLI program.
func (c *CLI) Run(args []string) {
	counts := make(map[string][]string)
	files := args[1:]
	if len(files) == 0 {
		countLines("-", c.inStream, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(c.errStream, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	var keys []string
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if len(counts[k]) > 1 {
			sort.Strings(counts[k])
			fmt.Fprintf(c.outStream, "%d\t%s\t%s\n",
				len(counts[k]), strings.Join(counts[k], ","), k)
		}
	}
}

func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	cli.Run(os.Args)
}

func countLines(filename string, f io.Reader, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], filename)
	}
}
