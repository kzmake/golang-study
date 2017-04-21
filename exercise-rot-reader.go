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

func Rot13(char, initChar byte) byte {
	return (char-initChar+13)%26 + initChar
}

func (r *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := r.r.Read(bytes)

	if err != nil {
		return 0, err
	}

	lowerCaseReg := regexp.MustCompile(`[a-z]`)
	upperCaseReg := regexp.MustCompile(`[A-Z]`)
	for i, c := range bytes {
		if lowerCaseReg.Match([]byte{c}) {
			bytes[i] = Rot13(c, 'a')
		} else if upperCaseReg.Match([]byte{c}) {
			bytes[i] = Rot13(c, 'A')
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
