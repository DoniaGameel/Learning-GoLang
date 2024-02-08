package main
/*
import "fmt"

func main(){
	//finite loop
	for i:=0; i<2; i++{
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

		formattedProfit := fmt.Sprintf("Profit: %.1f\n", profit)
		formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)
		fmt.Print(formattedProfit, formattedRatio)
	}

	// inifinite loop with break and continue
	for {
		// variables defenition
		var revenue float64
		var expenses float64
		var taxRate float64
		// taking input
		fmt.Print("Enter Revenue: ")
		fmt.Scan(&revenue)
		fmt.Print("Enter Expenses: ")
		fmt.Scan(&expenses)
		if expenses < 0 {
			continue
		}
		fmt.Print("Enter Tax Rate: ")
		fmt.Scan(&taxRate)
		if taxRate < 0 {
			break
		}
		// caculation
		var EBT = revenue - expenses
		var profit = EBT - EBT * taxRate / 100
		var ratio = EBT / profit
		//printing results
		fmt.Printf("EBT: %v\n",EBT)

		formattedProfit := fmt.Sprintf("Profit: %.1f\n", profit)
		formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)
		fmt.Print(formattedProfit, formattedRatio)
	}
}*/