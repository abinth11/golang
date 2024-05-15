package main

import "fmt"

func main() {
	capital := 80000
	if capital < 10000 {
		fmt.Println("Small cap user")
	} else if capital >= 10000 && capital <= 50000 {
		fmt.Println("Mid cap user")
	} else {
		fmt.Println("Large cap user")
	}

	if 11%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 5; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
}
