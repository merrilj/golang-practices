package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	merril := person{
		firstName: "Merril",
		lastName:  "Jeffs",
		contactInfo: contactInfo{
			email:   "jeffsmerril@gmail.com",
			zipCode: 80216,
		},
	}

	merril.updateName("Mer")
	merril.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}