package main

import "fmt"

func main() {
	c := add(1, 2)
	// fmt.Println(c+1) Go complains that the type is different.

	// using generics
	d := addGenrics(1, 3)
	fmt.Println(d + 1)
}

func add(a, b interface{}) interface{} { // func add[T int|float64|any](a, b T) T {

	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	return "unknown"
}

func addGenrics[T int | float64](a, b T) T {
	return a + b
}
