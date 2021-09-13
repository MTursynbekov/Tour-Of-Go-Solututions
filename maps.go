package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	word_count := make(map[string]int)
	for i, word := range words {
		cnt := 0
		for j := i; j < len(words); j++ {
			if word == words[j] {
				cnt += 1
			}
		}
		_, is_exists := word_count[word]
		if !is_exists {
			word_count[word] = cnt
		}
	}
	return word_count
}

func main() {
	wc.Test(WordCount)
}
