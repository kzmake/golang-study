package main

import (
	"io"
	"os"
	"regexp"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := r.r.Read(bytes)

	if err != nil {
		return 0, err
	}

	rot13 := func(c, init_c byte) byte {
		return (c-init_c+13)%26 + init_c
	}

	for i, c := range bytes {
		if regexp.MustCompile(`[a-z]`).Match([]byte{c}) {
			bytes[i] = rot13(c, 'a')
		} else if regexp.MustCompile(`[A-Z]`).Match([]byte{c}) {
			bytes[i] = rot13(c, 'A')
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
