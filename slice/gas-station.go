package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	currentFuel := 0
	startIndex := 0
	totalGas, totalCost := 0, 0
	for i, v := range gas {
		totalGas += v
		totalCost += cost[i]
		currentFuel += v - cost[i]
		fmt.Println(currentFuel)
		if currentFuel < 0 && i < len(gas)-1 {
			fmt.Println(currentFuel, cost[i])
			currentFuel = 0
			startIndex = i + 1
		}
	}
	if totalCost > totalGas {
		return -1
	}
	return startIndex
}
