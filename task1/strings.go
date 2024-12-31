package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	// expore strings package
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	s := "A man, a plan, a canal: Panama"
	s2 := s
	s = strings.Map(f, s)
	fmt.Println(s)
	s2 = strings.ToLower(s2)
	reg := regexp.MustCompile("[^a-z0-9]+")
	s2 = reg.ReplaceAllString(s2, "")
	fmt.Println(s2)
}
