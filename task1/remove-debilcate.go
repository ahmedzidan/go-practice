package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "geeksforgeeks"
	fmt.Println(removeDeblicate(s))
}

func removeDeblicate(s string) string {
	uniqueMap := map[rune]struct{}{}
	//uniqueString := ""
	var uniqueBuilde strings.Builder
	for _, ch := range s {
		if _, ok := uniqueMap[ch]; !ok {
			//uniqueString += string(ch)
			uniqueMap[ch] = struct{}{}
			uniqueBuilde.WriteRune(ch)
		}
	}
	return uniqueBuilde.String()
}
