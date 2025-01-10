package main

import "fmt"

type Notifier interface {
	notify()
}
type User struct {
	Name  string
	Email string
}

type Admin struct {
	User  // Embedding
	Level string
}

func (u *User) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.Name, u.Email)
}

func main() {
	ad := Admin{
		User: User{
			Name:  "john",
			Email: "test@test",
		},
		Level: "super",
	}
	ad.notify()
	sendNotfication(&ad)
}

func sendNotfication(n Notifier) {
	n.notify()
}
