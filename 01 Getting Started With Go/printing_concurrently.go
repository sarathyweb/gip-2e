package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 5)
	}
}

func main() {
	// Both the main() and count() runs simultaneously
	go count()
	time.Sleep(time.Millisecond * 20)
	fmt.Println("Hello world")
	time.Sleep(time.Millisecond * 5)
}
