package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) fullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) canDrive() bool {
	if p.Age > 20 {
		return true
	} else {
		return false
	}
}

func main() {
	person1 := Person{FirstName: "ahmed", LastName: "zidan", Age: 30}
	person2 := Person{FirstName: "Adam", LastName: "Zidan", Age: 1}

	fmt.Println(person1.fullName(), " can Drive ", person1.canDrive())
	fmt.Println(person2.fullName(), " Can Drive ", person2.canDrive())
}
