package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var output = flag.String("output", "jpeg", "output format (gif, jpeg or png)")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -output format < input > output\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch *output {
	case "gif":
		err = gif.Encode(os.Stdout, img, &gif.Options{})
	case "jpg", "jpeg":
		err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95})
	case "png":
		err = png.Encode(os.Stdout, img)
	default:
		log.Fatalf("Unsupported output image format: %s", *output)
	}

	if err != nil {
		log.Fatal(err)
	}
}
