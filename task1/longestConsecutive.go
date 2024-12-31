package main

import (
	"fmt"
)

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	/*
	 put it in a map
	 then check if the key of the map
	*/
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		numsMap[v] = struct{}{}
	}
	longest := 0
	for k := range numsMap {
		if _, ok := numsMap[k-1]; !ok { // 99
			length := 1
			for {
				if _, ok := numsMap[k+length]; ok {
					length++
					continue
				}
				longest = max(length, longest)
				break
			}
		}
	}
	return longest
}
