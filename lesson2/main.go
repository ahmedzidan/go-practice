package main

import "fmt"

func main() {
	src := "hello"
	fmt.Println("src:", &src)
	for i, v := range src {

		fmt.Println(string(v), i)
		src = "world"
		//if i == 1 {
		//	fmt.Println("Value is", string(v))
		//}
	}
	fmt.Println(src)
}
