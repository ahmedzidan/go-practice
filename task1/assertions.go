package main

import (
	"fmt"
)

type Company struct {
	Name string
}

func main() {
	var x interface{} = Company{"this is my company"}
	c1 := x.(Company)

	fmt.Println(c1)

	if c2, ok := x.(Company); ok {
		fmt.Println(c2, ok)
	}
}
