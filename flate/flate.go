package flate

import (
	goflate "compress/flate"
	"io"

	"github.com/go-zoox/compress/compressor"
)

// New creates a flate compressor.
func New(bufSize ...int) *compressor.Compressor {
	bufSizeX := 64
	if len(bufSize) > 0 {
		bufSizeX = bufSize[0]
	}

	return &compressor.Compressor{
		BufSize: bufSizeX,
		NewReader: func(r io.Reader) (io.ReadCloser, error) {
			closer := goflate.NewReader(r)
			return closer, nil
		},
		NewWriter: func(w io.Writer) io.WriteCloser {
			closer, err := goflate.NewWriter(w, goflate.BestCompression)
			if err != nil {
				panic(err)
			}
			return closer
		},
	}
}
