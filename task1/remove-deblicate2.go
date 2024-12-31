package main

import (
	"fmt"
)

func main() {
	num := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates3(num))
}

// num [1,2,2,3] == remove deplicate and return the len(array)

func removeDuplicates3(nums []int) int { //{1,1,1, 2,2,2}
	mymap := map[int]int{}
	k := 0
	for i, v := range nums {
		if mymap[v] < 2 {
			mymap[v] = mymap[v] + 1
			nums[k] = nums[i]
			k++
		}

	}
	fmt.Println(nums)
	return k
}
