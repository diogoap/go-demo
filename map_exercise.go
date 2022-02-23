package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	wordMap := make(map[string]int)

	words := strings.Split(s, " ")

	for _, value := range words {
		count := wordMap[value]

		if count == 0 {
			wordMap[value] = 1
		} else {
			wordMap[value] += count
		}
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
