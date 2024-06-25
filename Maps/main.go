package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
	}

	shapes := make(map[string]int)
	shapes["square"] = 4
	shapes["triangle"] = 3

	var metals = map[string]string{}
	metals["iron"] = "hematite"
	metals["aluminium"] = "bauxite"

	fmt.Println(colors)
	fmt.Println(metals)
	fmt.Println(shapes)

	// Delete the key from map
	delete(colors, "green")

	iterateMap(metals)
}

func iterateMap(m map[string]string) {
	// Iterate through map
	for key, value := range m {
		fmt.Println(key + " -- " + value)
	}
}
