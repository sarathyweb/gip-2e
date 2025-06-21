package main

import (
	"fmt"
	"time"
)

// A channel has a type
func printCount(c chan int) {
	num := 0
	for num >= 0 { // as long as the input is an positive integer
		num = <-c           // receives data from the channel
		fmt.Print(num, " ") // prints the data
	}
}

func main() {
	// create a channel
	c := make(chan int)
	a := []int{5, 4, 8, 6, 1, 7, 3, 0, 9, -1}
	// calls the printCount function
	go printCount(c)
	for _, v := range a {
		c <- v // sends data to the go channel
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
