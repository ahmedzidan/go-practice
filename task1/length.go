package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a"
	fmt.Println(lenthOfLastWord2(s))
}
func lengthOfLastWord(s string) int {
	words := strings.Fields(s) // ["hello", "world"]
	fmt.Println(len(words[len(words)-1]))
	return 0
}
func lenthOfLastWord2(s string) int {
	var finalLength int //"   hello   world   "
	for i := len(s) - 1; i >= 0; i-- {
		if finalLength == 0 && s[i] == ' ' {
			continue
		} else {
			if s[i] == ' ' {
				break
			}
			finalLength += 1
		}
	}
	return finalLength
}
