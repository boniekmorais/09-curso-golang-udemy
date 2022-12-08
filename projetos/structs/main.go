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

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v\n\n", alex)

	fmt.Printf("First Name: %s\n", alex.firstName)
	fmt.Printf("Last  Name: %s\n\n", alex.lastName)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Pointer operations

	// &variable => Give the memory address of the value.
	// *pointer  => Give the value this memory address is pointing at.

	// Using point getting address

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	// Shortcut without get the address

	jim.updateName("Jimmy")

	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v\n\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
