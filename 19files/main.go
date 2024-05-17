package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Creating a file and writing content to it
	fileContent := "Sample text"
	file, err := os.Create("example.txt")
	checkNilErr(err)

	// Writing content to the file
	length, err := io.WriteString(file, fileContent)
	checkNilErr(err)
	fmt.Println("Length of the file is", length)

	// Defer closing the file to ensure it's closed after main() exits
	defer file.Close()

	// Reading content from the file
	readFile("example.txt")
}

// Function to read content from a file and print it
func readFile(filename string) {
	// Reading content from the file
	contentByte, err := os.ReadFile(filename)
	checkNilErr(err)

	// Converting content bytes to string
	contentString := string(contentByte)

	// Printing the file content
	fmt.Println("File content from read method", contentString)
}

// Function to check for nil errors and handle them
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
