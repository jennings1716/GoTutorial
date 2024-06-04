package main

import "fmt"

func main() {

	addCard("Plus 4")
	addCard("Plus 2")
	addCard("Plus 14")

	uno := getDeck()
	fmt.Println(uno)

	deleteCardAt(1)

	uno = getDeck()
	fmt.Println(uno)

}
