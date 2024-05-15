package main

import "fmt"

func main() {
	var ptr1 *int
	fmt.Println("Ptr 1=", ptr1)

	randomNumber := 30
	ptr2 := &randomNumber
	fmt.Println("Ptr 2=", ptr2)
	fmt.Println("Val of Ptr 2=", *ptr2)
	*ptr2 = *ptr2 * *ptr2
	*ptr2 = *ptr2 + 1
	fmt.Println("Final val of Ptr 2=", *ptr2)

}
