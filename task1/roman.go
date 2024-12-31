package main

import "fmt"

var letterMaps = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	finalNumber := 0
	var currentValue, lastValue int
	for i := len(s) - 1; i >= 0; i-- {
		currentValue = letterMaps[s[i]] //1
		if currentValue < lastValue {   // 1< 5
			finalNumber -= currentValue // 5-1 =4
		} else {
			finalNumber += currentValue // 5
		}
		lastValue = currentValue //1
	}
	return finalNumber
}
