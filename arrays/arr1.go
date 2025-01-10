package main

import "fmt"

func main() {
	arr := []string{"2024", "2024", "2026"}
	for i, v := range arr {
		arr[1] = "2025"
		if i == 1 {
			fmt.Println("Value is", v)
		}
	}
}
