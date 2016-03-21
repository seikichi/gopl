package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ReadStdinError
)

type CLI struct {
	inStream  io.Reader
	outStream io.Writer
}

func (c *CLI) Run(args []string) int {
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", args[0])
		flagSet.PrintDefaults()
	}
	hashType := flagSet.String("type", "sha256", "Hash algorithm (sha256, sha384 or sha512)")

	if err := flagSet.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	bytes, err := ioutil.ReadAll(c.inStream)
	if err != nil {
		return ReadStdinError
	}

	switch *hashType {
	case "sha256":
		fmt.Fprintf(c.outStream, "%x\n", sha256.Sum256(bytes))
	case "sha384":
		fmt.Fprintf(c.outStream, "%x\n", sha512.Sum384(bytes))
	case "sha512":
		fmt.Fprintf(c.outStream, "%x\n", sha512.Sum512(bytes))
	default:
		flagSet.Usage()
		return ExitCodeParseFlagError
	}
	return ExitCodeOK
}

func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout}
	os.Exit(cli.Run(os.Args))
}
