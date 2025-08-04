package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "test@example.com",
			zipCode: 12345,
		},
	}
	(&jim).updateName("Jimmy")
	jim.print()
}

func (pp *person) updateName(newName string) {
	(*pp).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
