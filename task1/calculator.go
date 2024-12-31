package main

import "fmt"

func calculate(a int, b int) []float64 {
	// your code goes here
	var sum, diff, product, quotient float64
	sum = float64(a + b)
	diff = float64(a - b)
	product = float64(a * b)
	quotient = float64(a / b)
	results := []float64{sum, diff, product, quotient}
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
