package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user1 := user{firstName: firstName, lastName: lastName, birthDate: birthdate}
	// ... do something awesome with that gathered data!

	printUserData(&user1)

	user1.receiverPrintUserData()
}

func printUserData(p *user) {
	fmt.Println(p.firstName, p.lastName, p.birthDate)
}

func (p user) receiverPrintUserData() {
	fmt.Println(p.firstName, p.lastName, p.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
