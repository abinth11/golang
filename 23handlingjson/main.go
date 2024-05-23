package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting to encode JSON...")

	enCodeJson()
}

// user struct represents a user with fields for first name, last name,
// email, password, and favorite foods. The struct tags define how
// each field will be encoded into JSON.
//   - The `json:"firstName"` tag indicates that the FirstName field
//     will be encoded as "firstName" in JSON.
//   - The `json:"lastName"` tag indicates that the LastName field
//     will be encoded as "lastName" in JSON.
//   - The `json:"email"` tag indicates that the Email field
//     will be encoded as "email" in JSON.
//   - The `json:"-"` tag means the Password field will be ignored
//     and not included in the JSON output.
//   - The `json:"favoriteFoods,omitempty"` tag means the FavoriteFoods
//     field will be encoded as "favoriteFoods" in JSON, but will be
//     omitted if it is empty or nil.
type user struct {
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	Email         string   `json:"email"`
	Password      string   `json:"-"`
	FavoriteFoods []string `json:"favoriteFoods,omitempty"`
}

func enCodeJson() {

	users := []user{
		{
			FirstName:     "Alice",
			LastName:      "Johnson",
			Email:         "alice@example.com",
			Password:      "password123",
			FavoriteFoods: []string{"Pizza", "Sushi", "Ice Cream"},
		},
		{
			FirstName:     "Bob",
			LastName:      "Smith",
			Email:         "bob@example.com",
			Password:      "password456",
			FavoriteFoods: []string{"Burgers", "Steak", "Fries"},
		},
		{
			FirstName:     "Carol",
			LastName:      "Davis",
			Email:         "carol@example.com",
			Password:      "password789",
			FavoriteFoods: []string{"Pasta", "Salad", "Chocolate"},
		},
		{
			FirstName:     "David",
			LastName:      "Wilson",
			Email:         "david@example.com",
			Password:      "password101",
			FavoriteFoods: []string{"Sushi", "Ramen", "Tempura"},
		},
		{
			FirstName:     "Eve",
			LastName:      "Brown",
			Email:         "eve@example.com",
			Password:      "password202",
			FavoriteFoods: nil,
		},
	}

	// Marshal the slice of user structs into JSON with indentation for readability.
	// json.MarshalIndent takes three arguments:
	// 1. The data to be marshaled (in this case, the 'users' slice).
	// 2. A prefix string for each line (in this case, an empty string).
	// 3. An indentation string (in this case, two spaces) to format the output.
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalf("Error encoding JSON: %s", err)
	}

	// Marshal the slice of user structs into compact JSON.
	// json.Marshal takes one argument:
	// 1. The data to be marshaled (in this case, the 'users' slice).
	json, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Error encoding JSON: %s", err)
	}

	fmt.Printf("%s\n", json)

	fmt.Println(string(jsonData))
}
