package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}

	for _, word := range strings.Fields(s) {
		_, ok := m[word]

		if ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
