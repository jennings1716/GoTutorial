package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(`
	Hello
	This
	Is
	Jennings
	
	`)
	var playCards deck
	playCards = deck{}

	playCards = playCards.addCard("Plus 4")
	playCards = playCards.addCard("Pluse 2")
	playCards = playCards.addCard("Plus 14")

	fmt.Println(playCards.getDeck())

	// playCards.deleteCardAt(1)

	fmt.Println(playCards.toString())

	playCards.saveToFile("mycards.txt")
	fmt.Println("Read from file")
	fmt.Println(readFromFile("mycards.txt"))

	for index, _ := range playCards {
		swap := rand.Intn(3)
		fmt.Println(index)
		fmt.Println(swap)
		temp := playCards[index]
		playCards[index] = playCards[swap]
		playCards[swap] = temp
	}
	fmt.Println(playCards)

}
