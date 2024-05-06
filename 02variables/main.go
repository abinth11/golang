package main

import "fmt"

// * THIS IS A PUBLIC VARIABLE (FIRST LETTER CAPITAL, ACCESSIBLE IN ANY OTHER FILES)
const UserId string = "randomId"

func main() {

	//* String
	var username string = "benikyoo"
	fmt.Println("Username = ", username)
	fmt.Printf("Type of the variable: %T \n", username)

	//* Bool
	var isAdmin bool = true
	fmt.Println("isAdmin = ", isAdmin)
	fmt.Printf("Type of the variable: %T \n", isAdmin)

	//* Numbers
	var smallVal uint8 = 255
	fmt.Println("smallVal = ", smallVal)
	fmt.Printf("Type of the variable: %T \n", smallVal)

	var number int = 292939
	fmt.Println("number = ", number)
	fmt.Printf("Type of the variable: %T \n", number)

	var smallFloat float32 = 292939.912392393
	fmt.Println("smallFloat = ", smallFloat)
	fmt.Printf("Type of the variable: %T \n", smallFloat)

	var floatingPoint float64 = 292939.912392393
	fmt.Println("floatingPoint = ", floatingPoint)
	fmt.Printf("Type of the variable: %T \n", floatingPoint)

	//* DEFAULT VALUES AND ALIASES
	var randomNumVar int
	var randomStrVar string
	fmt.Println("randomNumVar =", randomNumVar)
	fmt.Println("randomStrVar =", randomStrVar)

	//* IMPLICIT TYPE
	var name = "John Doe"
	fmt.Println(name)

	/*
	* NO VAR STYLE
	* := VALOROUS OPERATOR
	* THIS OPERATOR IS ONLY ALLOWED INSIDE A METHOD
	 */
	sum := 10
	fmt.Println(sum)
	fmt.Println("MY userId= ", UserId)

}
