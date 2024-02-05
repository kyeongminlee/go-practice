package wordCount

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

	ch := make(chan string)

	go func() {
		defer close(ch)

		for _, word := range words {
			wg.Add(1)
			go func(word string) {
				defer wg.Done()
				ch <- word
			}(word)
		}

		wg.Wait()
	}()

	for word := range ch {
		wordMap[word]++
	}

	return wordMap
}
