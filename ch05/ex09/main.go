package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(expand(string(bs), func(s string) string { return s + s }))
}

var re = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	bs := re.ReplaceAllFunc([]byte(s), func(bs []byte) []byte {
		return []byte(f(string(bs[1:]))) // use [1:] to remove `$`
	})
	return string(bs)
}
