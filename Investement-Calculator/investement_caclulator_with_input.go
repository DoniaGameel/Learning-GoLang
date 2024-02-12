package main

/* Can be written as: 
import(
	"math"
	"fmt"
)*/
import "math"
import "fmt"

func main(){
	const inflationRate = 6.5
	var investementAmount float64 = 1000
	expectedReturnRate := 5.5
	var years = 10.0

	fmt.Scan(&investementAmount)
	
	var futureValue = investementAmount * math.Pow(1 + expectedReturnRate / 100, years)
	var futureRealValue = futureValue * math.Pow(1 + inflationRate / 100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}