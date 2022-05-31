package compress

import (
	"github.com/go-zoox/compress/compressor"
	"github.com/go-zoox/compress/flate"
	"github.com/go-zoox/compress/gzip"
	"github.com/go-zoox/compress/lzw"
	"github.com/go-zoox/compress/zlib"
)

// NewZlib creates a zlib compressor.
func NewZlib(bufSize ...int) *compressor.Compressor {
	return zlib.New(bufSize...)
}

// NewGzip creates a gzip compressor.
func NewGzip(bufSize ...int) *compressor.Compressor {
	return gzip.New(bufSize...)
}

// NewFlate creates a flate compressor.
func NewFlate(bufSize ...int) *compressor.Compressor {
	return flate.New(bufSize...)
}

// NewLZW creates a lzw compressor.
func NewLZW(bufSize ...int) *compressor.Compressor {
	return lzw.New(bufSize...)
}
