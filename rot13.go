package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'a' <= b && b <= 'm':
		return b + 13
	case 'A' <= b && b <= 'M':
		return b + 13
	case 'n' <= b && b <= 'z':
		return b - 13
	case 'N' <= b && b <= 'Z':
		return b - 13
	default:
		return b
	}
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
