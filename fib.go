//fib.go - Calculate and display first twenty Fibonacci Sequence Values
//  This program uses a re-entrant process to calculate the Fibonacci Sequence.
package main

import (
	"fmt"
)

//
// Calculate Fibonacci Number
//
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//
// Display first 20 numbers in Fibonacci Sequence
//
func main() {
	for n := 0; n < 20; n++ {

		fmt.Print(fib(n), " ") // Calculate & Display next value in the sequence
	}
	fmt.Println("")
}
