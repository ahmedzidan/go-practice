package main

import "fmt"

func main() {
	num := []int{1, 1, 2}
	fmt.Println(removeDuplicates(num))
}

// num [1,2,2,3] == remove deplicate and return the len(array)

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	finalMap := map[int]bool{}

	for _, v := range nums {
		if !finalMap[v] {
			finalMap[v] = true
		}
	}
	return len(finalMap)
}

func removeDuplicates(nums []int) int { //{1, 1, 2}
	if len(nums) <= 1 {
		return len(nums)
	}
	k := 0
	for i := range nums {
		if nums[i] != nums[k] {
			k += 1
			nums[k] = nums[i]
		}
	}
	return k + 1
}
