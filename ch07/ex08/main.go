package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type multiSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (s multiSort) Len() int {
	return len(s.t)
}

func (s multiSort) Less(i, j int) bool {
	return s.less(s.t[i], s.t[j])
}

func (s multiSort) Swap(i, j int) {
	s.t[i], s.t[j] = s.t[j], s.t[i]
}

func (s multiSort) byTitle() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byArtist() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byAlbum() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Album != y.Album {
			return x.Album < y.Album
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byYear() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byLength() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return s.less(x, y)
	}}
}

func newMultiSort(t []*Track) multiSort {
	return multiSort{t, func(_, _ *Track) bool { return false }}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	fmt.Println("Original:")
	printTracks(tracks)

	sort.Sort(newMultiSort(tracks).byYear().byTitle())
	fmt.Println("\nbyYear().byTitle():")
	printTracks(tracks)
}
