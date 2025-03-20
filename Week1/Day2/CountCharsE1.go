package main

import (
	"fmt"
	"strings"
	"sync"
)

func countLetters(str string, ch chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	letterCnt := make(map[string]int)
	str = strings.ToLower(str)
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			letterCnt[string(char)]++
		}
	}
	ch <- letterCnt
}

func main() {
	stringsList := []string{"quick", "brown", "fox", "lazy", "dog"}
	ch := make(chan map[string]int, len(stringsList))
	var wg sync.WaitGroup
	for _, str := range stringsList {
		wg.Add(1)
		go countLetters(str, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	finalFreq := make(map[string]int)

	for freqMap := range ch {
		for letter, count := range freqMap {
			finalFreq[letter] += count
		}
	}

	fmt.Println(finalFreq)
}
