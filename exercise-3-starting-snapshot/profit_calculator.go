package main

import "fmt"
import "io/ioutil"
import "errors"

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil{
		fmt.Print(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")

	if err != nil{
		panic(err)
	}
	taxRate,err := getUserInput("Tax Rate: ")

	if err != nil{
		fmt.Print(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	EBTText := fmt.Sprint(ebt)
	ProfitText := fmt.Sprint(profit)
	RatioText := fmt.Sprint(ratio)

	output := fmt.Sprint("EBT: %.1f\nProfit: %.1f\nRation: %.1f",ebt,profit,ratio)
	ioutil.WriteFile("balance.txt", []byte(output), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <=0 {
		return 0, errors.New("Can't be non-positive value")
	}
	return userInput, nil
}
