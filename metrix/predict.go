package main

import "fmt"

type DataFreq struct {
	word string
	Freq int
}

func prdict(s string, data map[string]string) string {
	return data[s]
}

func main() {
	data := [][]string{
		{"I", "am", "Sam"},
		{"Sam", "I", "like"},
		{"I", "like", "green"},
	}

	uniqueWord := processData(data)
	fmt.Println(prdict("I", uniqueWord))
}

func processData(data [][]string) map[string]string {
	uniqueWord := make(map[string]string)
	tempFreq := make(map[string][]DataFreq)

	for i, row := range data {
		for j, val := range row {
			if j+1 < len(row) {
				nextWord := data[i][j+1]
				if _, ok := tempFreq[val]; !ok {
					tempFreq[val] = []DataFreq{{word: nextWord, Freq: 1}}
					uniqueWord[val] = nextWord
				} else {
					found := false
					for idx, wordFreq := range tempFreq[val] {
						if wordFreq.word == nextWord {
							tempFreq[val][idx].Freq++
							found = true
						}
					}
					if !found {
						tempFreq[val] = append(tempFreq[val], DataFreq{word: nextWord, Freq: 1})
					}
					currentFreq := 1
					for _, highFreq := range tempFreq[val] {
						if highFreq.Freq > currentFreq {
							currentFreq = highFreq.Freq
							uniqueWord[val] = highFreq.word
						}
					}
				}
			} else {
				tempFreq[val] = append(tempFreq[val], DataFreq{word: "", Freq: 1})
				uniqueWord[val] = ""
			}
		}
	}
	return uniqueWord
}
