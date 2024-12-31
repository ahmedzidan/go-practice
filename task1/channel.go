package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "I"
		ch <- "love"
		ch <- "you."
	}()
	msg1 := <-ch
	msg2 := <-ch
	msg3 := <-ch
	fmt.Println(msg1, msg2, msg3)

}
