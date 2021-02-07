package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var amie person
	fmt.Println(amie)
	fmt.Printf("%+v\n", amie)

	amie.firstName = "Amie"
	amie.lastName = "Baker"

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// this doesn't change the jim's firstname
	jim.updateName("Jimmy")
	jim.print()

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}