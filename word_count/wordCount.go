package main

import "strings"

func WordCount(sentence string) map[string]int {
	words := strings.Split(sentence, " ")
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	return wordMap
}
