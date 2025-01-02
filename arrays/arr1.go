package main

import "fmt"

func main() {

	arr := [3]string{"2", "2", "2"}

	for i, v := range &arr {
		arr[1] = "1"
		if i == 1 {
			fmt.Println("Value is", v)
			fmt.Println("Value is", arr[i])
		}
	}
}
