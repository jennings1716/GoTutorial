package main

import "fmt"

func main() {

	for {
		fmt.Println("Welcome to Go Bank!!")
		fmt.Println("What do you want to do??")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		fmt.Printf("You chose %d \n", choice)

		if choice == 1 {
			fmt.Println("Check your Balance")
		} else if choice == 2 {
			fmt.Println("Deposit Money")
		} else if choice == 3 {
			fmt.Println("Money Withdrawn")
		} else {
			fmt.Println("Exit")
			break
		}
	}
	fmt.Println("Thanks for choosing our Bank!!")
}
