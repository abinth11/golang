package main

import "fmt"

// * LOOPS
func main() {
	/*
		Notes:
		1. For Loop (Traditional):
		   This is the traditional for loop commonly used in many programming languages.
		   It uses an index variable to iterate over a range.
		   Example:
		   for i := 0; i < len(users); i++ {
		       fmt.Println(users[i])
		   }

		2. For Loop (Range with Index):
		   This loop uses the range keyword to iterate over the elements of a collection.
		   It retrieves the index of each element.
		   Example:
		   for i := range users {
		       fmt.Println(users[i])
		   }

		3. For Loop (Range with Value):
		   This loop also uses the range keyword but retrieves the value of each element directly.
		   It ignores the index.
		   Example:
		   for _, user := range users {
		       fmt.Println(user)
		   }

		4. For Loop (While-like):
		   This loop resembles a while loop and continues as long as the condition is true.
		   It uses `continue` to skip the current iteration and `goto` to jump to a labeled statement.
		   Example:
		   randomNum := 1
		   for randomNum < 15 {
		       if randomNum == 5 {
		           goto lco
		       }
		       if randomNum == 4 {
		           randomNum++
		           continue
		       }
		       fmt.Println("Value is", randomNum)
		       randomNum++
		   }

		5. Goto Statement:
		   The `goto` statement transfers control to a labeled statement within the same function.
		   It can be used to exit loops or jump to specific parts of the code.
		   Example:
		   lco:
		       fmt.Println("jumping to... => goto")
	*/

	// var users = []string{"John doe", "Patrick", "Sasi", "Dane", "Mike"}

	//* type one
	// for i := 0; i < len(users); i++ {
	// 	fmt.Println(users[i])
	// }

	//* type two
	// for i := range users {
	// 	fmt.Println(users[i])
	// }

	//* type three
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	//* type four => similar to while loop
	randomNum := 1
	for randomNum < 15 {
		// if randomNum == 3 {
		// 	break
		// }
		if randomNum == 5 {
			goto lco
		}
		if randomNum == 4 {
			randomNum++
			continue
		}
		fmt.Println("Value is", randomNum)
		randomNum++
	}

	//* GOTO STATEMENT , can use any name for it like a variable
lco:
	fmt.Println("jumping to... => goto")
}
