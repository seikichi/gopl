package zip

import (
	"archive/zip"
	"errors"
	"io"
	"os"

	"github.com/seikichi/gopl/ch10/ex02/archive"
)

type reader struct {
	r       *zip.Reader
	i       int
	current io.ReadCloser
}

func (r *reader) Next() (*archive.Header, error) {
	r.i++
	if r.i >= len(r.r.File) {
		return nil, io.EOF
	}

	if r.current != nil {
		r.current.Close()
	}

	f := r.r.File[r.i]
	reader, err := f.Open()
	if err != nil {
		return nil, err
	}

	r.current = reader
	return &archive.Header{Name: f.Name}, nil
}

func (r *reader) Read(b []byte) (n int, err error) {
	if r.current == nil {
		return 0, errors.New("Invalid use of zip reader, use Next()")
	}
	return r.current.Read(b)
}

func read(f *os.File) (archive.Reader, error) {
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	r, err := zip.NewReader(f, info.Size())
	if err != nil {
		return nil, err
	}

	return &reader{r: r, i: -1}, nil
}

func init() {
	archive.RegisterFormat("zip", "\x50\x4b", 0, read)
}
