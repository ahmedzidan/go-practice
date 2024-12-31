package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotat2(num, k)
}

func rotate(nums []int, k int) {
	temp := make([]int, len(nums))
	for k > 0 {
		LastValue := nums[len(nums)-1]
		for j := 1; j < len(nums); j++ {
			temp[j] = nums[j-1]
		}
		temp[0] = LastValue
		copy(nums, temp)
		k--
	}
	fmt.Println(nums)
}

func rotat2(nums []int, k int) {
	temp := make([]int, len(nums))
	//3 +0 % 7
	for i := 0; i < len(nums); i++ {
		indx := (i + k) % len(nums)
		temp[indx] = nums[i]
	}
	copy(nums, temp)
}
