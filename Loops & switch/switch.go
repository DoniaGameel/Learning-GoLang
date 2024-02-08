package main

import "fmt"

func main(){
	var choice int
	fmt.Println("Enter choice: ")
	fmt.Scan(&choice)

	switch choice{
	case 1:
		fmt.Println("Best choice")
	case 2:
		fmt.Println("Correct! But still not the best")
	default:
		fmt.Println("Wrong!")
	}
}