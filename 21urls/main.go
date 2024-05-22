package main

import (
	"fmt"
	"net/url"
)

// Define a constant for the API URL to be parsed and manipulated.
const apiUrl string = "https://example.com:3000/user?search=shoes&limit=10&skip=2"

func main() {
	fmt.Println(apiUrl)

	// PARSING A URL
	// Parse the URL string into a URL structure using url.Parse.
	result, err := url.Parse(apiUrl)
	if err != nil {
		// Handle any error that occurs during parsing.
		panic(err)
	}

	// Access and print the scheme (e.g., "https").
	fmt.Println(result.Scheme)
	// Access and print the host (e.g., "example.com:3000").
	fmt.Println(result.Host)
	// Access and print the path (e.g., "/user").
	fmt.Println(result.Path)
	// Access and print the port (e.g., "3000").
	fmt.Println(result.Port())
	// Access and print the raw query string (e.g., "search=shoes&limit=10&skip=2").
	fmt.Println(result.RawQuery)

	// Query parameters are accessed using the Query method, which returns a map of query parameters.
	queryParams := result.Query()
	// Print the type of query parameters (should be url.Values).
	fmt.Printf("Type of query params is %T\n", queryParams)
	// Access and print the value of the "search" parameter.
	fmt.Println(queryParams["search"])

	// Loop through and print each value in the query parameters.
	for _, val := range queryParams {
		fmt.Println(val)
	}

	// Construct a URL from its parts using the url.URL struct.
	partsOfUrl := &url.URL{
		Scheme:  "https",       // Define the scheme (e.g., "https").
		Host:    "example.com", // Define the host (e.g., "example.com").
		Path:    "/example",    // Define the path (e.g., "/example").
		RawPath: "search=abcd", // Define the raw path/query string (e.g., "search=abcd").
	}
	// Build the URL string from the given parts.
	urlFromParts := partsOfUrl.String()
	fmt.Println(urlFromParts)
}
