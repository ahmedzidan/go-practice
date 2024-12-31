package main

import "fmt"

func main() {

	array1 := []int{}
	m := 0
	array2 := []int{1}
	n := 1

	merge(array1, m, array2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 { //len = m + n
		nums1 = nums2[0:n]
	} else {
		for j := len(nums1) - 1; j > 0; j-- {
			if n == 0 {
				break
			}
			if nums1[m-1] > nums2[n-1] {
				nums1[j] = nums1[m-1]
				m--
			} else {
				nums1[j] = nums2[n-1]
				n--
			}
		}
	}
	//array1 = [1,2,3,0,0] m=3   array2= [2,5,6] n=3

	fmt.Println(nums1)
}
