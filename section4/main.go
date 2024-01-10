package main

import (
	"fmt"
)

// @L43
type person struct {
	firsName string
	lastName string
	contact  contactInfo
	contactInfo
}

// @46
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	contact := contactInfo{email: "randomEmail@email.kk", zipCode: 80000}
	// @L44
	alex := person{"Alex", "Anderson", contact, contactInfo{email: "randomEmail@mail.kk", zipCode: 8888}}
	fmt.Println(alex)

	johh := person{firsName: "John", lastName: "Mac"}
	fmt.Println(johh.firsName, johh.lastName)

	// @L45
	var mac person
	mac.firsName = "MacFirst"
	mac.lastName = "MacLast"
	fmt.Println(mac)
	mac.print()

	// @L46
	jim := person{
		firsName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updatedName("Hello")
	jim.print()

	// @L49
	// jimPointer := &jim
	// jimPointer.updatedNamePointer("Pointer")
	// jimPointer.print()

	// @L51 shortcut version of L49
	jim.updatedNamePointer("JimPointer")
	jim.print()

}

// @L47
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updatedName(newFirstName string) {
	p.firsName = newFirstName
}

// @L49 / @L50
// *person this a type description, it measns we're working with a pointe to a person
// *pointerToPerson, this is an operator - it means we want to manipulate the value the pointer is referenci
func (pointerToPerson *person) updatedNamePointer(newFirstName string) {
	(*pointerToPerson).firsName = newFirstName
}
