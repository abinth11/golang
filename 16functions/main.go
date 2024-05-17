package main

import "fmt"

func main() {
	// Calling functions and printing their results
	printMessage()
	strResponse := getMessage()
	fmt.Println(strResponse)
	result := sum(3, 5)
	fmt.Println("Sum is =", result)
	nSum, resultString, thirdResult := sumUptoN(2, 323, 23, 23, 232, 323, 2323)
	fmt.Println("Sum of n numbers =", nSum, resultString, thirdResult)
}

// Normal Function: Prints a message to the console.
func printMessage() {
	fmt.Println("Hi from function")
}

// Function with Return Value: Returns a string message.
func getMessage() string {
	return "From get message function"
}

// Function with Parameters: Accepts two integers and returns their sum.
func sum(num1 int, num2 int) int {
	return num1 + num2
}

// Variadic Function: Accepts a variable number of integers and returns their sum along with two additional strings.
func sumUptoN(values ...int) (int, string, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum, "Can return more than one value", "Third value"
}
