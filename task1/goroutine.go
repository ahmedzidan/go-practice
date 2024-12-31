package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func(ms string) {
		defer wg.Done()
		fmt.Println(ms)
	}("test go routine1")
	wg.Add(1)
	go func(ms string) {
		defer wg.Done()
		fmt.Println(ms)
	}("test 2")
	wg.Wait()
}
