package main

import "fmt"

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

	jim := person{
		firstName: "Kunal",
		lastName:  "Dhingra",
		contactInfo: contactInfo{
			email:   "kdhingra55@gmail.com",
			zipCode: 94000,
		},
	}

	jim.print()
	jim.updateName("Kanu")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
