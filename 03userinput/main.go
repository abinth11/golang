package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "Hii User"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")

	//* COMMA OK SYNTAX || ERROR OK SYNTAX
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered input =", input)
	fmt.Printf("Type of the input = %T", input)
}
