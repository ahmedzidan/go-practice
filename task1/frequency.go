package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	frequencyMap := map[string]int{}
	words := strings.Fields(text)
	for _, word := range words {
		if _, ok := frequencyMap[word]; ok {
			frequencyMap[word] = frequencyMap[word] + 1
		} else {
			frequencyMap[word] = 1
		}
	}
	return frequencyMap
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))
}
