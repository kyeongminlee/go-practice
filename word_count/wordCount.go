package main

import (
	"strings"
	"sync"
)

var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func WordCount(sentence string) map[string]int {
	words := strings.Split(sentence, " ")
	wordMap := make(map[string]int)

	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()

			mutex.Lock()
			defer mutex.Unlock()
			wordMap[word]++
		}(word)
	}

	wg.Wait()

	return wordMap
}
