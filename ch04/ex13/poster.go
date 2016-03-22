package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	Poster string
}

func main() {
	movie := getMovieInfo(os.Args[1:])
	downloadPoster(movie)
}

func getMovieInfo(query []string) Movie {
	req, _ := http.NewRequest("GET", "https://omdbapi.com", nil)
	q := req.URL.Query()
	q.Add("t", strings.Join(query, "+"))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("HTTP Get failed: %s", err)
	}
	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	return movie
}

func downloadPoster(movie Movie) {
	if movie.Poster == "" {
		log.Fatalf("Poster URL is invalid: %q", movie.Poster)
	}

	resp, err := http.Get(movie.Poster)
	if err != nil {
		log.Fatalf("Poster download failed: %s", err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("Poster output to stdout failed: %s", err)
	}
}
