package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name
	fmt.Println("The variable is name and it stores a string 'bill' ")
	fmt.Println("Location where actually string 'bill' is stored ", &namePointer) // address in which name pointer is stored

	printAddress(namePointer)
}

func printAddress(namePointer *string) {
	fmt.Println("Location where actually string copy of 'bill' address is stored ", namePointer)
	fmt.Println(*namePointer)
}
