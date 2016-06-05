package main

import "testing"

type testWriter struct {
	writeArg []byte
}

func (w *testWriter) Write(p []byte) (int, error) {
	w.writeArg = append([]byte(nil), p...)
	return len(p), nil
}

func TestCountingWriter(t *testing.T) {
	original := &testWriter{}
	wrapped, counter := CountingWriter(original)

	bs := []byte("Hello, world!")
	wrapped.Write(bs)

	if *counter != int64(len(bs)) {
		t.Errorf("After wrapped.Write(%q), counter becomes %v; want %v", bs, *counter, int64(len(bs)))
	}

	if string(original.writeArg) != string(bs) {
		t.Errorf("Wrapped writer.Write did not called correctly")
	}
}
