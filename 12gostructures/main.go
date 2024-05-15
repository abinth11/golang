package main

import "fmt"

// * STRUCTURES(SIMILAR TO CLASSES)
/*
 * <===== NOTES =====>
 * No concept of inheritance , super or parent
 * Use first letter upper case to make it as public
 *
 */
func main() {
	john := User{"John Doe", "john@gmail.com", false, 20}
	fmt.Println("Struct", john)
	fmt.Println("Name", john.Name)
	fmt.Printf("Name = %v, Email= %v, Age= %v\n", john.Name, john.Email, john.Age)

	//with +v placeholder , will display in more detail
	fmt.Printf("Detailed view %+v", john)
}

type User struct {
	Name       string
	Email      string
	IsLoggedIn bool
	Age        int
}
