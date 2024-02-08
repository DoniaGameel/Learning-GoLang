package main

import "fmt"

func main(){
	outputText("AB", "CD")
	x,y := returnTwoFloats()
	fmt.Println(x,y)

	w,z := returnTwoFloats_2()
	fmt.Println(w,z)
}

//void function
func outputText(text1,text2 string){
	fmt.Println(text1)
	fmt.Println(text2)
}

// function with return
func returnTwoFloats() (float64, float64){
	x := 1.0
	y := 2.0
	return x, y
}

// Alternative return value syntax
func returnTwoFloats_2() (x float64,y float64){
	x = 3.0
	y = 4.0
	//also works
	//return x, y
	return
}