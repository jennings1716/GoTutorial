package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 3, 4, 6}
	fmt.Println(transformNumbers(&numbers, double))

	tetraTransform := createTransformer(4)
	fmt.Println(tetraTransform(10))
}

func transformNumbers(numbers *[]int, transform transformFn) ([]int, []int) {
	doubledNumbers := []int{}
	tripleddNumbers := []int{}

	for _, value := range *numbers {
		doubledNumbers = append(doubledNumbers, transform(value))
	}

	for _, value := range *numbers {
		tripleddNumbers = append(tripleddNumbers, (func(num int) int {
			return num * 3
		})(value))
	}
	return doubledNumbers, tripleddNumbers
}

func double(num int) int {
	return num * 2
}

func createTransformer(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
