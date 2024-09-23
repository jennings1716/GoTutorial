package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user1, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	user.PrintUserData(user1)
	// user1.receiverPrintUserData()
	user1.ClearUserName()
	user1.ReceiverPrintUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
