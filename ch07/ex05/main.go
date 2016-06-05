package main

import "io"

type limitedReader struct {
	r io.Reader
	n int64
}

func (r *limitedReader) Read(p []byte) (int, error) {
	if r.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.n {
		p = p[0:r.n]
	}
	n, err := r.r.Read(p)
	r.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{r, n}
}

func main() {}
