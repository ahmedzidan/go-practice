package main

import "fmt"

func main() {
	num1 := []int{1, 0}
	m := 1
	num2 := []int{1}
	n := 1
	mergeArr(num1, m, num2, n)
	// m+n is the lenth of the new array
	// 1 < 2 ? yes [1]
	// 2 < 2? no [1,2]
	// 2 < 5? yes [1,2,2]
	// 3< 5 ? yes [1,2,2,3]
	// [1,2,2,3,5,6]
}
func mergeArr(num1 []int, m int, num2 []int, n int) {
	num2Index := 0
	num1Index := 0
	finalArray := make([]int, m+n)
	if m == 0 {
		copy(num1, num2)
	} else if n > 0 {
		for i := 0; i < m+n; i++ {
			if num1Index < len(num1) && num1[num1Index] < num2[num2Index] && m != num1Index {
				finalArray[i] = num1[num1Index]
				num1Index += 1
			} else if n != num2Index {
				finalArray[i] = num2[num2Index]
				num2Index += 1
			}
			if n == num2Index {
				fmt.Println(finalArray)
				fmt.Println(num1[num1Index:m])
				for num1Index < m {
					finalArray[i+1] = num1[num1Index]
					i++
					num1Index++
				}
				break
			}
			// Append slice2 to slice1

		}
		copy(num1, finalArray)
	}
	fmt.Println(num1)
}
