package main

import (
	"fmt"
	wc "word_count/wordCount"
)

func main() {
	text := "apple banana apple orange kiwi"
	wordMap := wc.WordCount(text)
	fmt.Println(wordMap)
}
