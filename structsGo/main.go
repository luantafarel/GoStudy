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
	people := person{firstName: "Luan", lastName: "Tafarel", contact: contactInfo{
		email:   "luan@gmail.com",
		zipCode: 304959,
	}}
	people.print()
	people.updateName("Tafarel da Silva")
	people.print()
}

func (p *person) updateName(newLastName string) {
	p.lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
