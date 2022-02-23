package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this rot13Reader) Read(b []byte) (n int, err error) {

	n, err = this.r.Read(b)

	for i := 0; i < n; i++ {
		c := b[i]

		if (c >= 'A' && c <= 'M') || (c >= 'a' && c <= 'm') {
			b[i] = c + 13
		} else {
			b[i] = c - 13
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
