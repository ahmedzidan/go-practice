package main

import "fmt"

func main() {
	num := []int{3, 3, 4}
	fmt.Println(majorityElement(num))
}

func majorityElement(nums []int) int {
	myMap := map[int]int{}
	major := 0
	Majorcount := 0
	for _, v := range nums {
		myMap[v] = myMap[v] + 1
		if myMap[v] > Majorcount {
			major = v
			Majorcount = myMap[v]
		}
	}
	return major
}
