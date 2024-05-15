package main

import "fmt"

func main() {
	user := make(map[string]string)
	user["Name"] = "John Doe"
	user["Age"] = "20"
	user["Favorite Food"] = "Pizza"
	user["Pet"] = "Dog"
	fmt.Println(user)

	//how to get single val
	fmt.Println("Name= ", user["Name"])

	//how to delete a value
	delete(user, "Age")
	fmt.Println("User after deleting", user)

	//how to loop a map
	for key, val := range user {
		fmt.Printf("Key is= %v,& Value is= %v \n", key, val)
	}

	//if value is not required user underscore, like comma ok syntax
	for _, val := range user {
		fmt.Printf("Value is= %v \n", val)
	}
}
