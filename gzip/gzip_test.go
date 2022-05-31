package gzip

import "testing"

func TestGzip(t *testing.T) {
	c := New()

	src := "Hello, World!"
	dst := c.Compress([]byte(src))

	if decoded, err := c.Decompress(dst); err != nil {
		t.Error(err)
	} else if string(decoded) != src {
		t.Errorf("%s != %s", string(decoded), src)
	}
}
