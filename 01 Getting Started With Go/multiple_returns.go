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
}
