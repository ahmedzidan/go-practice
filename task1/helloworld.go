package main

import "fmt"

func main() {
	var name1 string
	var name2 = "test"
	name1 = "test2"
	name3 := name2 + " " + name1
	fmt.Println("hello zidan")
	fmt.Println(name3)

	var isBig bool

	isBig = false

	println(isBig)

	switch {
	case isBig == true:
		fmt.Println("here is true")
	case !isBig:
		fmt.Println("I will kill you")
	default:
		fmt.Println("Iam the default bye")
	}

	var array1 []int
	array2 := []int{1, 2, 3, 4, 5}

	for _, value := range array2 {
		fmt.Println(value)
	}
	array1 = array2[1:4]
	fmt.Println(array1)
	copyslice := make([]int, len(array2))

	copy(copyslice, array2)

	fmt.Println(copyslice)
}
