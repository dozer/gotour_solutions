package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

// Find the count of each word in a string
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
