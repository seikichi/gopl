package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Package struct {
	Dir  string
	Name string
	Deps []string
}

func main() {
	checkArguments()

	packageNames, err := getPackageNames(os.Args[1:])
	if err != nil {
		log.Fatalf("Failed to get packages names: %s", err)
	}

	dependents, err := getPackagesDependOn(packageNames)
	if err != nil {
		log.Fatalf("Failed to get dependents: %s", err)
	}

	for _, p := range dependents {
		fmt.Printf("%s (in %s)\n", p.Name, p.Dir)
	}
}

func checkArguments() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s package\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func getPackageNames(packages []string) ([]string, error) {
	args := append([]string{"list", "-e", "-json"}, packages...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}

	packageNames := []string{}
	d := json.NewDecoder(bytes.NewReader(out))
	for {
		var p Package
		err := d.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		packageNames = append(packageNames, p.Name)
	}
	return packageNames, nil
}

func getPackagesDependOn(packageNames []string) ([]Package, error) {
	packages := []Package{}

	out, err := exec.Command("go", "list", "-e", "-json", "...").Output()
	if err != nil {
		return nil, err
	}
	d := json.NewDecoder(bytes.NewReader(out))

loop:
	for {
		var p Package
		err := d.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, name := range packageNames {
			found := false
			for _, dep := range p.Deps {
				if name == dep {
					found = true
					break
				}
			}
			if !found {
				continue loop
			}
		}

		packages = append(packages, p)
	}
	return packages, nil
}
