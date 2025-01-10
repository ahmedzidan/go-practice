package main

import "fmt"

func main() {
	//binary tree

	// How to get the matching characters and non-matching characters in a string.
	//
	//Example:-
	//Input: str1 = “abcdef”, str2 = “defghia”
	//Output: 4
	//Matching characters are: a, d, e, f
	//Similarly find the non-matching charcters for str1 and str2.
	//
	fmt.Println(matchingString("abcdef", "defghia"))
}

func matchingString(str1 string, str2 string) int {
	sortedStr := make(map[string]int)
	finalMatch := 0
	for _, v := range str1 {
		sortedStr[string(v)] = 1 // a =. 1
	}
	for _, v := range str2 {
		if _, ok := sortedStr[string(v)]; ok {
			finalMatch += 1
		}
	}
	return finalMatch
}

//ou have 3 bikes with 20L of fuel tank capacity and 100kmpl of mileage.
// All the bikes are at same starting point.
// You have to travel maximum distance and carry the bikes until the entire fuel is burned from a bike.
//How you can maximise the distance.
//Start with 3 driver and leave the bike and driver once the bike fuel tank is empty.
//If required make some assumptions.

// 3 bikes => longest distance
// 20L fuel tank capacity
// 1l = 100km
// start point = 0 * 100km = 0
// 20l = 2000km  p1, p2, p3
// p1 = 20l p2= 19 p3 = 18
// p1 = 20l p2 = 18 p3 = 16  p3= 0
// p1 = 20/3 = 7 20-7 = p2= 19.5 p1 = 19.5 p3 = 19/2 = 9
// 19.5 % 2 = 9

func maxDistance() {
	TotalFuelTank := 20
	mileage := 100
	distance := 0
	for i := 3; i > 0; i-- {
		currentFuelTank := TotalFuelTank - TotalFuelTank%i // 20%3 = 13
		distance += (TotalFuelTank % i) * mileage          // 7*100 = 700
		TotalFuelTank = TotalFuelTank%i + currentFuelTank  // 20/3 = 6 +13 = 19
	}
	fmt.Println(distance)
}
