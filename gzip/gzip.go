package gzip

import (
	gogzip "compress/gzip"
	"io"

	"github.com/go-zoox/compress/compressor"
)

// New creates a gzip compressor.
func New(bufSize ...int) *compressor.Compressor {
	bufSizeX := 64
	if len(bufSize) > 0 {
		bufSizeX = bufSize[0]
	}

	return &compressor.Compressor{
		BufSize: bufSizeX,
		NewReader: func(r io.Reader) (io.ReadCloser, error) {
			return gogzip.NewReader(r)
		},
		NewWriter: func(w io.Writer) io.WriteCloser {
			return gogzip.NewWriter(w)
		},
	}
}
