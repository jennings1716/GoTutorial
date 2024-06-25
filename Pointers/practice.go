package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer) // address in which name pointer is stored

	printAddress(namePointer)
}

func printAddress(namePointer *string) {
	fmt.Println(&namePointer)
}
