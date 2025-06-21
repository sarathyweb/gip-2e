package main

import "fmt"

// Nameed return values allows use to automatically initalize the return variables
func getSrings() (name1 string, name2 string) {
	name1 = "Sarathy"
	name2 = "Duck"
	return
}

func main() {
	n1, n2 := getSrings()
	fmt.Println("Name 1 is ", n1, "\nName 2 is ", n2)
}
