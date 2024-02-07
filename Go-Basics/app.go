// package: must be the first statement in any file
// you split your code across packages to organize your code
// we can import features from one package to another
// 'main' is a special package name ==> the main entry point
// if main is removed and we build ==> no file as an output as no entry point is found
package main 

// fmt is a built in package
import "fmt"

func main () {
	// "Hello world!" is called 'value'
	fmt.Print("Hello world!")
}
