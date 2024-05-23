package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	const apiBaseUrl string = "http://localhost:8080"
	const getRequestUrl string = apiBaseUrl + "/get-message"
	const postRequestUrl string = apiBaseUrl + "/save-user"
	const postFormUlr string = apiBaseUrl + "/handle-post-form"
	handleGetRequest(getRequestUrl)
	handlePostRequest(postRequestUrl)
	handlePostFormRequest(postFormUlr)
}

// handleGetRequest sends an HTTP GET request to the specified URL,
// reads the response body, and prints the status code and response body.
func handleGetRequest(baseUrl string) {
	// Send an HTTP GET request to the specified URL
	response, err := http.Get(baseUrl)
	// Check for errors during the request
	if err != nil {
		// If an error occurs, panic and terminate the program
		panic(err)
	}
	// Ensure the response body is closed after the function returns
	defer response.Body.Close()

	fmt.Println("HTTP Status:", response.StatusCode)

	// Check if the response status code indicates success (200 OK)
	if response.StatusCode == http.StatusOK {
		fmt.Println("Request was successful!")
		// Read the response body into a byte slice
		content, err := io.ReadAll(response.Body)
		// Check for errors while reading the response body
		if err != nil {
			// If an error occurs, print it and return
			fmt.Println(err)
			return
		}

		// Create a strings.Builder to efficiently build a string
		var responseString strings.Builder
		// Write the content of the response to the strings.Builder
		byteCount, err2 := responseString.Write(content)
		// Check for errors while writing to the strings.Builder
		if err2 != nil {
			// If an error occurs, print it and return
			fmt.Println(err2)
			return
		}

		fmt.Println("Byte count is:", byteCount)
		fmt.Println("Response Body:", responseString.String())
	} else {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
	}
}

func handlePostRequest(baseUrl string) {
	requestBody := strings.NewReader(`{
		"name":"Jane Doe",
		"age":"22",
		"email":"jane@gmail.com"
	}`)

	response, err := http.Post(baseUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("Successful request")
		responseContent, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(responseContent))
	} else {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)

	}

}

func handlePostFormRequest(baseUrl string) {

	body := url.Values{}
	body.Add("username", "aizen233")
	body.Add("password", "2323")

	response, err := http.PostForm(baseUrl, body)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("Successful request")
		responseContent, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(responseContent))
	} else {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)

	}
}
