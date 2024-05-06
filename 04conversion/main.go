package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	message := "Please Enter a number"
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println("You are entered", input)

	if err != nil {
		fmt.Println("Error occurred while reading the input")
	} else {
		parsedInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("Error occurred while parsing the string", err)
		} else {
			fmt.Println("Sum =", parsedInput+10)
		}
	}

}
