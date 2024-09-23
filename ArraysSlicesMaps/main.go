package main

import "fmt"

type anymap map[any]any

type product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"test", "test2"}
	prices := [4]float64{10.1, 12.3, 11.00, 22.32}
	fmt.Print(prices[2])
	fmt.Print(productNames)

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	// dynamic prices
	dynamic_prices := []float64{12.2, 34.2}
	dynamic_prices[0] = 1
	new_dynamic_price := append(dynamic_prices, 45.22)
	fmt.Println(new_dynamic_price)

	discount := []float64{1.2, 3.4, 4.5}
	new_discount := append(new_dynamic_price, discount...)
	fmt.Println(new_discount)

	make_array := make([]string, 2, 5)
	fmt.Println(cap(make_array))
	fmt.Println(len(make_array))
	fmt.Println(make_array[0])
	make_array[0] = "12"
	fmt.Println(make_array)

	//-------------------------------------------------------------------------
	// MAPs

	data_types := anymap{"int": 1, "string": "test", "float": 2.3}

	fmt.Println(data_types)

	data_types["struct"] = product{title: "testbook", id: "123", price: 12}

	data_types["string"] = "testing"
	fmt.Println(data_types)

	//-------------------------------------------------------------------------
	// Looping Maps and Arrays

	for key, value := range data_types {
		fmt.Println("key ", key)
		fmt.Println("value ", value)
	}
}
