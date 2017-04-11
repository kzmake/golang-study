package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}

	for _, word := range strings.Fields(s) {
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
