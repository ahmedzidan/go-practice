package main

import "fmt"

func main() {
	num := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(num, val))
}
func removeElement(num []int, val int) int {
	k := 0
	for _, value := range num {
		if value != val {
			num[k] = value
			k++
		}
	}
	return k
}
