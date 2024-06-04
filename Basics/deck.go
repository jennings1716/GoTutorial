package main

import "fmt"

// create a new type deck

type deck []string

var playingCards deck = deck{}

func addCard(card string) {
	playingCards = append(playingCards, card)
	fmt.Println("Added a new card")
}

func getDeck() deck {
	return playingCards
}

func deleteCardAt(index int) deck {
	return append(playingCards[:index], playingCards[index+1:]...)
}
