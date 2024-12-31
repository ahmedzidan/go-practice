package main

import "fmt"

func main() {
	nums := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println(largestSum(nums))
	//s := "babad"
	//fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	// babad
	// start = b  end = d ? no
	// start = a  end = a == yes
	// aba
	if len(s) == 0 {
		return ""
	}
	// odd ..i..
	// even ..i,i+1,..
	maxLength := 1
	var start, end int
	for i := 0; i < len(s); i++ {
		start1, end1, len1 := expandAroundCenter(s, i, i) // babad, 0, 0
		start2, end2, len2 := expandAroundCenter(s, i, i+1)

		if len1 > maxLength {
			start = start1
			end = end1
			maxLength = len1
		}
		if len2 > maxLength {
			start = start2
			end = end2
			maxLength = len2
		}
	}
	return s[start : end+1]
}
func expandAroundCenter(s string, left, right int) (int, int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] { //babad, -1 >=0 yes, 0 < 5 yes, b = b ? yes
		left--
		right++
	}
	start := left + 1
	end := right - 1
	length := end - start + 1
	return start, end, length
}

func largestSum(nums []int) int {
	var maxSoFar int
	totalSum := 0
	start, end, s := 0, 0, 0
	// -5
	for i := 0; i < len(nums); i++ {
		totalSum = totalSum + nums[i]
		if maxSoFar < totalSum {
			maxSoFar = totalSum
			start = s
			end = i
		}
		if totalSum < 0 {
			totalSum = 0
			s++
		}
	}
	fmt.Println(nums[start : end+1])
	//
	// -2 -3 = -5
	// -2 -3 4 = -1
	// -2 -3 4 -1 = -2
	// -2 -3 4 -1 -2 = -3
	// -2 -3 4 -1 -2 1 -2
	// -2 5 = 3 -2 5    index = max and also sum = max
	// 0
	return maxSoFar
}
