package main

import "fmt"

// Go allows a function to return more than one value
func getSrings() (string, string) {
	return "Foo", "Bar"
}

func main() {

	// The first retuned value will be stored in s1
	// the second returned value is stored in s2
	s1, s2 := getSrings()
	fmt.Println("Value of s1 ", s1, "\nValue of s2 ", s2)

	// if we don't want to omit any of the values, we can use the _ operator
	//1. ignore first returned values
	_, p2 := getSrings()
	fmt.Println("Value of p2 ", p2)

	//2. ignore second parameter
	p1, _ := getSrings()
	fmt.Println("The value of p1 ", p1)

	//3. ignore both first and second paramter
	// NOTE that we are not using the := operator as we are not creating new variables
	_, _ = getSrings() // which is same as calling getSrings()
	fmt.Println("We have nothing to print")
}
