package main

import "fmt"

func main() {
    age := 32
    var agePointer *int = &age
    fmt.Println(age)
    adultYears := getAdultYears(agePointer)
    fmt.Println(adultYears)
}

func getAdultYears(age *int) int { // Corrected return type to int
    return *age - 18
}
