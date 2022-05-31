package compressor

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Compressor is a common interface for all compressions.
type Compressor struct {
	BufSize   int
	NewWriter func(io.Writer) io.WriteCloser
	NewReader func(io.Reader) (io.ReadCloser, error)
}

// Compress encodes the raw bytes.
func (c *Compressor) Compress(src []byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, c.BufSize))
	w := c.NewWriter(buf)
	w.Write(src)
	w.Close()

	return buf.Bytes()
}

// Decompress decodes the encoded bytes.
func (c *Compressor) Decompress(src []byte) ([]byte, error) {
	r, err := c.NewReader(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}

	dst, err := ioutil.ReadAll(r)
	r.Close()

	return dst, err
}
