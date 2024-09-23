package main

import "fmt"

type product struct {
	name  string
	price float64
}

func main() {
	hobbies := []string{"movies", "running", "jogging"}

	fmt.Println(hobbies)

	fmt.Println("First data ", hobbies[0])

	least := hobbies[1:]

	fmt.Println("Least data ", least)

	priority := hobbies[:2]

	fmt.Println("Priority data ", priority)
	fmt.Println("Priority size ", cap(priority))

	changePriority := priority[1:3]
	fmt.Println(changePriority)

	goals := []string{"GoExpert", "AWSExpert"}
	fmt.Println(goals)

	products := []product{{name: "pen", price: 10.01}, {name: "pencil", price: 20.01}}

	fmt.Println(products)

}
