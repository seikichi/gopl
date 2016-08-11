package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/seikichi/gopl/ch10/ex02/archive"
	_ "github.com/seikichi/gopl/ch10/ex02/archive/tar"
	_ "github.com/seikichi/gopl/ch10/ex02/archive/zip"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s archive output-directory\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	r, _, err := archive.Read(file)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(os.Stderr, "Extract %s ...\n", h.Name)

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, r); err != nil {
			log.Fatalln(err)
		}

		output := os.Args[2]

		p := filepath.Join(output, h.Name)
		d, _ := filepath.Split(p)

		if _, err = os.Stat(d); err != nil {
			os.MkdirAll(d, 0755)
		}

		if err = ioutil.WriteFile(p, buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
		}
	}
}
