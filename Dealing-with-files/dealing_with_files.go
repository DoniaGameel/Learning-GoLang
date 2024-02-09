package main

import "fmt"
import "io/ioutil"
import "strconv"
import "errors"

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
	fmt.Println("EBT",EBT)
	fmt.Println("Profit",profit)
	fmt.Println("Ratio",ratio)
	EBTText := fmt.Sprint(EBT)
	//writing to a file
	//0644: the file permission
	ioutil.WriteFile("balance.txt", []byte(EBTText), 0644)
	// reading from a file
	// when we don't find the file ==> return empty byte collection whhich will be comverted to empty string
	data, err:=ioutil.ReadFile("balance.txt")
	if err != nil{
		return 1000, errors.New("Failed to read file.")
	}
	balanceTxt := string(data)
	balance, _ := strconv.ParseFloat(balanceTxt,64)
	fmt.Println(balance)
}