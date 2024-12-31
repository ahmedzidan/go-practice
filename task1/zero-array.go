package main

import (
	"fmt"
)

func main() {
	nums := []int{6, 3, -1, -3, 4, -2, 2, 4, 6, -12, -7}
	findZeroArrayUsingMap(nums)
}
func printSubArray(nums []int) [][]int {
	var currentSum = nums[0]
	start, end := 0, 1
	result := [][]int{}
	for start < end && end < len(nums) {
		currentSum += nums[end]
		if currentSum == 0 {
			subArrasy := nums[start : end+1]
			result = append(result, subArrasy)
			end++
		} else {
			end++
		}
		if end == len(nums)-1 {
			start++
			end = start + 1
			currentSum = nums[start]
		}
	}
	return result
}
func findZeroArrayUsingMap(nums []int) {
	sumMap := make(map[int][]int)
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 {
			fmt.Println("index from 0 to ", i)
		}

		if indecies, ok := sumMap[sum]; ok {
			fmt.Println(sumMap)
			fmt.Println(sumMap[sum], "i = ", i, sum)
			for _, index := range indecies {
				fmt.Printf("from index %d to index %d\n", index+1, i)
			}
		}
		sumMap[sum] = append(sumMap[sum], i)
	}

}
