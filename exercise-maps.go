package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)

	for _, word := range strings.Fields(s) {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
