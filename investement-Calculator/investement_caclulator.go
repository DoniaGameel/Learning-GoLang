package main

/* Can be written as: 
import(
	"math"
	"fmt"
)
import "math"
import "fmt"

func main(){
	// both are float64
	var x, y float64 = 1000, 10
	// also works
	w, z := 1000, 10
	// type annotation
	//similar to: var investementAmount = 1000
	var investementAmount float64 = 1000
	// it means the type should be inferred by Go
	// Can't add float64 ==> error
	expectedReturnRate := 5.5
	var years = 10
	const inflationRate = 6.5
	var futureValue = float64(investementAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
	var futureRealValue = futureValue * math.Pow(1 + inflationRate / 100, float64(years))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(w)
	fmt.Println(z)
	
}
*/