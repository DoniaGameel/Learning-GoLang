 package main

 import "fmt"
/*
 func main(){
	//[key_type]value_type
	websites := map[string]string{ // without a blank
		"Google": "https://google.com",
		"Amazon web services": "https://aws.com",
	} 

	fmt.Println(websites)
	// accesing element of the map
	fmt.Println(websites["Google"])
	// adding element to the map
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	// deleting element from the map
	delete(websites, "Google")
	fmt.Println(websites)
}
*/

// type aliases
type floatMap map[string]float64

func (m floatMap) output(){
	fmt.PrintlN(m)
}

func main(){
	// 3 pre allocated elements
	//courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	// already pre allocated locations
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	// here will re allocate
	courseRatings["flutter"] = 4.7
	fmt.Println(courseRatings)

	for key, value := range courseRatings{
		fmt.Println(key)
		fmt.Println(value)
	}
}