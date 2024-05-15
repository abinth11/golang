package main

import "fmt"

func main() {
	var students [3]string
	students[0] = "Student 1"
	students[2] = "Student 3"
	fmt.Println("Second element", students[1])
	fmt.Println("Students", students)
	fmt.Println("Length of the array", len(students))

	var users = [3]string{"User 1", "User 2", "User 3"}
	fmt.Println("Users", users)
}
