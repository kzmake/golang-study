package main

import (
	"io"
	"os"
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

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
