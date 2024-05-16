package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		Notes:
		1. No `break` Needed:
		   In Go, each `case` in a `switch` statement automatically breaks after executing,
		   meaning it doesn't fall through to the next case unless explicitly stated with
		   a `fallthrough` keyword. This is different from languages like C, C++, or Java,
		   where you must use `break` to prevent fall-through.

		2. Use of `fallthrough`:
		   The `fallthrough` keyword in Go is used to explicitly state that the execution
		   should continue to the next case, regardless of the case's condition. This can be
		   useful when you want multiple cases to execute in sequence for certain conditions.
		   In the provided code, using `fallthrough` after `case 1` and `case 2` means that
		   when `diesValue` is 1 or 2, the subsequent `case` statements will also execute sequentially.
	*/

	diesValue := rand.Intn(6) + 1
	switch diesValue {
	case 1:
		fmt.Println("Value is 1")
		fallthrough
	case 2:
		fmt.Println("Value is 2")
		fallthrough
	case 3:
		fmt.Println("Value is 3")
	case 4:
		fmt.Println("Value is 4")
	case 5:
		fmt.Println("Value is 5")
	case 6:
		fmt.Println("Value is 6")
	default:
		fmt.Println("No valid value found")
	}
}
