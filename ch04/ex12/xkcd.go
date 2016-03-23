package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var usage = `Usage xkcd <command> [options]

Commands:
  download: Download xkcd comics to local
  index   : Create index file from downloaded comics
  search  : Search comics by keyword
`

func main() {
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	switch args[0] {
	case "download":
		download(args[1:])
	case "index":
		index(args[1:])
	case "search":
		search(args[1:])
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func download(args []string) {
	fs := flag.NewFlagSet("download", flag.ExitOnError)
	var directory string
	var from int
	fs.StringVar(&directory, "directory", ".", "download directory")
	fs.IntVar(&from, "from", 1, "first download comic id, useful to resume download")
	fs.Parse(args)

	for i := from; ; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		log.Printf("download %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			os.Exit(0)
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			os.Exit(0)
		}

		filename := filepath.Join(directory, fmt.Sprintf("%d.json", i))
		ioutil.WriteFile(filename, b, os.ModePerm)
	}
}

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

type Index struct {
	Docs map[string]map[int]bool
}

func index(args []string) {
	fs := flag.NewFlagSet("index", flag.ExitOnError)
	var directory, out string
	fs.StringVar(&directory, "directory", ".", "download comic directory")
	fs.StringVar(&out, "out", "index.gov", "outputindex file")
	fs.Parse(args)

	matches, err := filepath.Glob(filepath.Join(directory, "*.json"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	index := Index{}
	index.Docs = map[string]map[int]bool{}

	for _, filename := range matches {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		var comic Comic
		json.Unmarshal(file, &comic)

		r := regexp.MustCompile(`[ {}\[\]<>:\!,."';?]`)
		transcript := r.ReplaceAllLiteralString(comic.Transcript, " ")
		s := bufio.NewScanner(strings.NewReader(transcript))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			if _, ok := index.Docs[s.Text()]; !ok {
				index.Docs[s.Text()] = map[int]bool{}
			}
			index.Docs[s.Text()][comic.Num] = true
		}
	}

	buf := new(bytes.Buffer)
	e := gob.NewEncoder(buf)
	if err := e.Encode(index); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(out, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func search(args []string) {
	fs := flag.NewFlagSet("search", flag.ExitOnError)
	var directory, out string
	fs.StringVar(&directory, "directory", ".", "download comic directory")
	fs.StringVar(&out, "index", "index.gov", "index file")
	fs.Parse(args)

	b, err := ioutil.ReadFile(out)
	if err != nil {
		log.Fatal(err)
	}

	var index Index
	d := gob.NewDecoder(bytes.NewReader(b))
	if err := d.Decode(&index); err != nil {
		log.Fatal(err)
	}

	result := map[int]bool{}
	for i, w := range fs.Args() {
		if i == 0 {
			result = index.Docs[w]
			continue
		}

		docs := index.Docs[w]
		for i := range result {
			if !docs[i] {
				delete(result, i)
			}
		}
	}

	for i := range result {
		filename := filepath.Join(directory, fmt.Sprintf("%d.json", i))
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		var comic Comic
		json.Unmarshal(file, &comic)
		fmt.Println(comic.Num, comic.Title)
	}
}
