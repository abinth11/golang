package main

import "fmt"

func main() {
	// Defer statements are executed in Last-In-First-Out (LIFO) order.
	// This means the last defer statement will be executed first, and so on.

	// The following defer statements will be executed first in LIFO order.
	defer fmt.Println("World") // Executed third
	defer fmt.Println("One")   // Executed second
	defer fmt.Println("\nTwo") // Executed first

	// This statement will be executed immediately.
	fmt.Println("Hello") // Executed fourth

	// Calling a function with defer statements
	myDefer() // Executed fifth
}

// Function with defer statements
func myDefer() {
	// In a loop, defer statements are evaluated immediately but executed after the surrounding function returns.
	for i := 0; i < 5; i++ {
		// Each defer statement is pushed onto a stack and executed in LIFO order.
		defer fmt.Print(i) // The value of 'i' at the time of the defer statement execution is printed.
	}
	// The following statement will be executed before any deferred statements.
	fmt.Println("\nEnd of myDefer function") // Executed sixth
}
