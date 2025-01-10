package main

import "fmt"

type Notifier interface {
	notify()
}

type User struct {
	Name  string
	Email string
}

func (u *User) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.Name, u.Email)
}

func sendNotification(n Notifier) {
	n.notify()
}

func main() {
	user := User{"Bill", "zidan@gmail.com"}
	sendNotification(&user)
}
