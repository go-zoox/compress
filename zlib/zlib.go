package zlib

import (
	gozlib "compress/zlib"
	"io"

	"github.com/go-zoox/compress/compressor"
)

// New creates a zlib compressor.
func New(bufSize ...int) *compressor.Compressor {
	bufSizeX := 64
	if len(bufSize) > 0 {
		bufSizeX = bufSize[0]
	}

	return &compressor.Compressor{
		BufSize: bufSizeX,
		NewReader: func(r io.Reader) (io.ReadCloser, error) {
			return gozlib.NewReader(r)
		},
		NewWriter: func(w io.Writer) io.WriteCloser {
			return gozlib.NewWriter(w)
		},
	}
}
