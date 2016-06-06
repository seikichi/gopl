package main

import (
	"sort"
	"testing"
)

func newData() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func BenchmarkStableSortRepeat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := newData()
		sort.Stable(customSort{s, func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			return false
		}})
		sort.Stable(customSort{s, func(x, y *Track) bool {
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
			return false
		}})
		sort.Stable(customSort{s, func(x, y *Track) bool {
			if x.Album != y.Album {
				return x.Album < y.Album
			}
			return false
		}})
		sort.Stable(customSort{s, func(x, y *Track) bool {
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			return false
		}})
	}
}

func BenchmarkMultiSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := newMultiSort(newData()).byTitle().byArtist().byAlbum().byYear()
		sort.Stable(s)
	}
}
