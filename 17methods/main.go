package main

import "fmt"

func main() {
	// Creating a User instance
	john := User{"John Doe", "john@gmail.com", false, 20}

	// Printing detailed view of the user
	fmt.Printf("Detailed view %+v\n", john)

	// Calling methods on the User instance
	john.GetIsLoggedIn() // Method without pointers
	john.AddOneToAge()   // Method without pointers

	// Printing detailed view of the user after method calls
	fmt.Printf("Detailed view %+v\n", john)

	// Calling methods with pointers
	// Method with pointers: This method updates the IsLoggedIn field of the User instance.
	john.SetIsLoggedIn(true)
	fmt.Printf("Detailed view %+v\n", john)

	// Method with pointers: This method increments the age of the User instance.
	john.AddOneToAgePointer()
	fmt.Printf("Detailed view %+v\n", john)
}

/* Define the User struct */
type User struct {
	Name       string
	Email      string
	IsLoggedIn bool
	Age        int
}

/*
  - Method without pointers:

This passes a copy of the struct. If the actual value needs to be changed,
pass a reference to the struct using a pointer.
*/
func (u User) GetIsLoggedIn() {
	fmt.Println("User logged in =", u.IsLoggedIn)
}

/*
  - Method without pointers:

This method adds one to the age of the User instance.
However, it operates on a copy of the struct, so the original struct remains unchanged.
*/
func (u User) AddOneToAge() {
	u.Age += 1
	fmt.Println("Age after adding one =", u.Age)
}

/*
  - Method with pointers:

This method updates the IsLoggedIn field of the User instance.
It operates on the original struct using a pointer, so changes are reflected outside the method.
*/
func (u *User) SetIsLoggedIn(loggedIn bool) {
	u.IsLoggedIn = loggedIn
	fmt.Println("User logged in status updated =", u.IsLoggedIn)
}

/*
  - Method with pointers:

This method increments the age of the User instance.
It operates on the original struct using a pointer, so changes are reflected outside the method.
*/
func (u *User) AddOneToAgePointer() {
	u.Age += 1
	fmt.Println("Age after adding one =", u.Age)
}
