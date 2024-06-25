package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type deck

type deck []string

func (d deck) addCard(card string) deck {
	d = append(d, card)
	fmt.Println(d)
	fmt.Println("Added a new card")
	return d
}

func (d deck) getDeck() deck {
	return d
}

func (d deck) deleteCardAt(index int) deck {
	return append(d[:index], d[index+1:]...)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	slicestr := string(bs)
	strarr := strings.Split(slicestr, ",")
	return deck(strarr)

}
