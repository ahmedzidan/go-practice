package main

import "fmt"

func main() {
	num1 := []int{1, 2, 3, 4, 6}
	num2 := []int{2, 4, 6, 8}
	fmt.Println(interWith2Pointers(num1, num2))
}
func intersection(num1, num2 []int) []int {
	intersection := make([]int, 0)
	//[1,2,3,4,6]
	// [2,4,6,8] == print all the number that are shown in both arrays.
	num1Map := make(map[int]struct{})
	for _, num := range num1 {
		num1Map[num] = struct{}{}
	}

	for _, num := range num2 {
		if _, ok := num1Map[num]; ok {
			intersection = append(intersection, num)
		}
	}
	return intersection
}

func interWith2Pointers(num1, num2 []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(num1) && j < len(num2) {
		if num1[i] < num2[j] {
			i++
		} else if num1[i] > num2[j] {
			j++
		} else {
			result = append(result, num2[j])
			i++
			j++
		}
	}
	return result
}
