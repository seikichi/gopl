package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type diskUsage struct {
	nfiles, nbytes int64
}

type work struct {
	root     string
	fileSize int64
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	usages := map[string]*diskUsage{}
	for _, root := range roots {
		usages[root] = &diskUsage{nfiles: 0, nbytes: 0}
	}

	works := make(chan work)
	var n sync.WaitGroup

	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, works)
	}

	go func() {
		n.Wait()
		close(works)
	}()

	tick := time.Tick(500 * time.Millisecond)
loop:
	for {
		select {
		case <-done:
			for _ = range works {
			}
			return
		case w, ok := <-works:
			if !ok {
				break loop
			}
			usages[w.root].nfiles++
			usages[w.root].nbytes += w.fileSize
		case <-tick:
			printDiskUsage(usages)
		}
	}
	printDiskUsage(usages)
}

func printDiskUsage(usages map[string]*diskUsage) {
	keys := []string{}
	for k := range usages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println()
	for _, k := range keys {
		fmt.Printf("%s %d files  %.1f GB\n", k, usages[k].nfiles, float64(usages[k].nbytes)/1e9)
	}
}

func walkDir(root string, dir string, n *sync.WaitGroup, works chan<- work) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, works)
		} else {
			works <- work{root: root, fileSize: entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}
