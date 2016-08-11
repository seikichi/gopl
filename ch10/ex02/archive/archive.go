package archive

import (
	"bufio"
	"errors"
	"os"
)

type Header struct {
	Name string
}

type Reader interface {
	Next() (*Header, error)
	Read(b []byte) (n int, err error)
}

func Read(file *os.File) (Reader, string, error) {
	f := sniff(file)
	file.Seek(0, os.SEEK_SET)
	if f.read == nil {
		return nil, "", errors.New("Invalid format")
	}
	r, err := f.read(file)
	return r, f.name, err
}

func RegisterFormat(name, magic string, offset int, read func(*os.File) (Reader, error)) {
	formats = append(formats, format{name, magic, offset, read})
}

type format struct {
	name, magic string
	offset      int
	read        func(*os.File) (Reader, error)
}

var formats []format

func sniff(file *os.File) format {
	r := bufio.NewReader(file)
	for _, f := range formats {
		b, err := r.Peek(f.offset + len(f.magic))
		if err == nil && match(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}
