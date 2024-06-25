package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := make([]byte, 100)
	arguments := os.Args
	file, err := os.Open(arguments[1]) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(data))
}
