package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{2, 2, 3, 3, 4}
	val := 3
	k := removeElement(nums, val)
	fmt.Println("lenth", k)
	for i := 0; i < k; i++ {
		println(nums[i])
	}
	
	fmt.Println("nums", nums)
}
func removeDuplicates(nums []int) int { //{1, 1, 2}
	if len(nums) <= 1 {
		return len(nums)
	}
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k += 1
			nums[k] = nums[i]
		}
	}
	return k + 1
}
