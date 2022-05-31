package lzw

import (
	golzw "compress/lzw"
	"io"

	"github.com/go-zoox/compress/compressor"
)

// New creates a lzw compressor.
func New(bufSize ...int) *compressor.Compressor {
	bufSizeX := 64
	if len(bufSize) > 0 {
		bufSizeX = bufSize[0]
	}

	return &compressor.Compressor{
		BufSize: bufSizeX,
		NewReader: func(r io.Reader) (io.ReadCloser, error) {
			closer := golzw.NewReader(r, golzw.LSB, 8)
			return closer, nil
		},
		NewWriter: func(w io.Writer) io.WriteCloser {
			return golzw.NewWriter(w, golzw.LSB, 8)
		},
	}
}
