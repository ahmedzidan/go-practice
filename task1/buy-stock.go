package main

import "fmt"

type PriceDay struct {
	Price int
	day   int
}

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		fmt.Println(prices[i], " min ", minPrice)
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			// If diff of current stock with minPrice is greater
			// update the profit
			profit = prices[i] - minPrice
		}
	}

	return profit
}
