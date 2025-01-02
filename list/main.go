package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(4)
	intList.PushBack(5)
	intList.PushFront(3)

	for element := intList.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)

	}
}
