package main

import "fmt"

func main() {
	//ratings := []int{29, 51, 87, 87, 72, 12} // expected 12
	//ratings := []int{1, 3, 2, 2, 1} // expected 7
	ratings := []int{1, 3, 4, 5, 2} // expected 11
	//ratings := []int{1, 2, 87, 87, 87, 2, 1} // expected 13
	fmt.Println(candy2(ratings))
}

func candy(ratings []int) int {
	candyNumber := 0
	var previousRate int
	previousCandy := 1

	for i := 0; i < len(ratings); i++ {
		checked := false
		if i >= 1 && ratings[i] > previousRate { //2 >
			fmt.Println("comper pre ", ratings[i], " i= ", i)
			candyNumber += previousCandy // 3
			previousCandy = previousCandy + 1
			checked = true
		}
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			if !checked && i > 0 && previousCandy > 1 {
				fmt.Println("comper next ", ratings[i], " i= ", i)
				candyNumber += previousCandy - 1 // 3+1 =4
				previousCandy = previousCandy - 1
				checked = true
			} else if !checked && previousCandy == 1 {
				fmt.Println("comper next ", ratings[i], " i= ", i)
				candyNumber += previousCandy + 1 // 3+1 =4
				previousCandy = previousCandy + 1
				checked = true
			}
		}
		if !checked {
			previousCandy = 1
		}
		candyNumber += 1
		fmt.Println("cany number", candyNumber, " previous candy ", previousCandy)
		previousRate = ratings[i] //2
	}
	return candyNumber
}

func candy2(rating []int) int {
	numOfCandy := map[int]int{}
	for i := 0; i < len(rating); i++ {
		numOfCandy[i] = 1
	}

	for ind := 1; ind < len(rating); ind++ {
		if rating[ind] > rating[ind-1] {
			numOfCandy[ind] = numOfCandy[ind-1] + 1
		}
	}

	for lastIndex := len(rating) - 2; lastIndex >= 0; lastIndex-- {
		if rating[lastIndex] > rating[lastIndex+1] && numOfCandy[lastIndex] <= numOfCandy[lastIndex+1] {
			numOfCandy[lastIndex] = numOfCandy[lastIndex+1] + 1
		}
	}
	TotalCanday := 0
	fmt.Println(numOfCandy)
	for _, val := range numOfCandy {
		TotalCanday += val
	}
	return TotalCanday
}
