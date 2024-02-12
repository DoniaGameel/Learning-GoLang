package main

import "fmt"
func main(){
	// Building dynamic lists with slices
	// starting values
	prices := []float64 {10.99 , 9.99}
	updatedPrices := append(prices, 5.99)
	fmt.Println(prices)
	fmt.Println(updatedPrices)
}
/*
func main(){
	// creating arrays
	prices := [4] float64{10.99 , 20.99, 40.00, 14.99}
	fmt.Println(prices)

	// define array wiith empty slots:
	var productNames [4] string
	// first item is assigned and the 3 other items are empty
	productNames = [4] string{ "A book" }
	fmt.Println(productNames)
	// output the first element
	fmt.Println(productNames[0])
	// et e valaue to an element
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	// slices
	// first is incuded, second is execluded
	featuredPrice := prices[1:3]
	fmt.Println(featuredPrice)

	// modifying the slice will modify the original array
	// slice is only a window in the original array
	featuredPrice[0] = 9.99
	fmt.Println(prices)

}*/