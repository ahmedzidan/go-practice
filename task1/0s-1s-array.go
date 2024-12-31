package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 0, 1, 2}
	fmt.Println(sort012Pointer(nums))
}
func sort012(nums []int) []int {
	//[0,2,1,0,1,2}
	//
	numsMap := map[int]int{}
	for _, v := range nums {
		numsMap[v] += 1
	}
	numZero := numsMap[0]
	numOne := numsMap[1]
	numTwo := numsMap[2]
	finalArray := []int{}
	for numZero != 0 {
		finalArray = append(finalArray, 0)
		numZero--
	}
	for numOne != 0 {
		finalArray = append(finalArray, 1)
		numOne--
	}
	for numTwo != 0 {
		finalArray = append(finalArray, 2)
		numTwo--
	}
	return finalArray
}

func sort012Pointer(nums []int) []int {
	// {0,2,0,1,2}
	lo, mid := 0, 0
	high := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[mid] == 0 {
			nums[lo], nums[mid] = nums[mid], nums[lo]
			mid++
			lo++
		} else if nums[mid] == 2 {
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		} else {
			mid++
		}
		fmt.Println(nums)
	}

	return nums
}
