package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(i byte) byte {
	switch {
	case 'a' <= i && i <= 'm' || 'A' <= i && i <= 'M':
		return i + 13
	case 'n' <= i && i <= 'z' || 'N' <= i && i <= 'Z':
		return i - 13
	default:
		return i
	}
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
