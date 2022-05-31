# compress
> Simply compressor, support zlib/gzip/flate/lzw.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/compress)](https://pkg.go.dev/github.com/go-zoox/compress)
[![Build Status](https://github.com/go-zoox/compress/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/compress/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/compress)](https://goreportcard.com/report/github.com/go-zoox/compress)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/compress/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/compress?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/compress.svg)](https://github.com/go-zoox/compress/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/compress.svg?label=Release)](https://github.com/go-zoox/compress/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/compress
```

## Getting Started

```go
package main

import (
	"fmt"

	"github.com/go-zoox/compress"
)

func main() {
	gzip := compress.NewGzip()
	bytes := gzip.Compress([]byte("hello world"))
	bytes2, _ := gzip.Decompress(bytes)
	fmt.Println(string(bytes2))
}
```

Try in [Go Playground](https://go.dev/play/p/4ATlBB8BHQi).

## Inspired By
* [cosiner/gohper](https://github.com/cosiner/gohper/blob/master/encoding/encoding.go) - common libs here

## License
GoZoox is released under the [MIT License](./LICENSE).
