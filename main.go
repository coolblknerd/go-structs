package main

import "fmt"

// Person is a struct
type Person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p Person) printName() {
	fmt.Printf("Hello %s %s", p.firstName, p.lastName)
	fmt.Printf("%+v", p)
}

// This is a type description - it means we're working with a pointer to a person
func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	person := Person{
		firstName: "Reggie",
		lastName:  "Davis",
		contactInfo: contactInfo{
			email: "reggied@email.com",
			zip:   88708,
		},
	}

	// personPointer := &pointer --> this points to the memory address
	person.updateName("Jordan")
	person.printName()
}

/*
When you create a new variable in Go, the runtime makes a copy in a different memory space

Turn an address into a value with *address
Turn value into address with &value
*/
