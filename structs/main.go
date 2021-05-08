package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main()  {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy")

	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}