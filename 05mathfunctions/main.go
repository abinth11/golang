package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math functions ")

	//* addition of int + float
	num1 := 4
	num2 := 5.9
	sum := num1 + int(num2)
	fmt.Println("Sum = ", sum)

	//* RANDOM NUMBERS => from math/rand
	// fmt.Println(rand.Intn(5))

	//* RANDOM NUMBERS => from crypto/rand
	randomVal, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("Random value = ", randomVal)
}
