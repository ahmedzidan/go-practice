package main

import (
	"fmt"
	"sort"
)

func main() {

	intervals := [][]int{{1, 2}, {8, 10}, {2, 6}, {15, 18}}
	fmt.Println(merge2(intervals))

}

func merge2(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	var finalInterval [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//{{1, 2}, {2, 6}, {8, 10}, {15, 18}}
	//{1,6}
	for i := 0; i < len(intervals); i++ {
		if i < len(intervals)-2 && intervals[i][1] >= intervals[i+1][0] {
			finalInterval = append(finalInterval, []int{intervals[i][0], intervals[i+1][1]})
			i++
		} else {
			finalInterval = append(finalInterval, intervals[i])
		}
	}

	return finalInterval
}
