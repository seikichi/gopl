package tar

import (
	"archive/tar"
	"os"

	"github.com/seikichi/gopl/ch10/ex02/archive"
)

type reader struct {
	r *tar.Reader
}

func (r *reader) Next() (*archive.Header, error) {
	h, err := r.r.Next()
	if err != nil {
		return nil, err
	}
	return &archive.Header{Name: h.Name}, nil
}

func (r *reader) Read(b []byte) (n int, err error) {
	return r.r.Read(b)
}

func read(f *os.File) (archive.Reader, error) {
	r := tar.NewReader(f)
	return &reader{r}, nil
}

func init() {
	archive.RegisterFormat("zip", "\x75\x73\x74\x61\x72", 0x101, read)
}
