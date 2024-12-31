package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["dog"] = "dog1"
	map1["dog2"] = "dog2"
	fmt.Println(map1["dog"])

	map2 := map[int]string{1: "test", 2: "test2"}

	for k, v := range map2 {
		fmt.Println(k, "  ", v)
	}
}
