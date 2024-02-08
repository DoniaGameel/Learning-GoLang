package main

import "fmt"

func main(){
	// variables defenition
	var revenue float64
	var expenses float64
	var taxRate float64
	// taking input
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)
	// caculation
	var EBT = revenue - expenses
	var profit = EBT - EBT * taxRate / 100
	var ratio = EBT / profit
	//printing results
	fmt.Printf("EBT: %v\n",EBT)
	fmt.Println("Profit",profit)
	fmt.Println("Ratio",ratio)
}