package main

import (
	"fmt"
	"io"
	"net/http"
)

// WEB REQUESTS
const url = "https://example.com"

func main() {
	// Sending an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	/*
	   The response from an HTTP request is an instance of the http.Response struct.
	   It contains several important fields, including:
	   - Status: The HTTP status code (e.g., 200 OK).
	   - StatusCode: The numeric status code (e.g., 200).
	   - Body: The response body, which is the content returned by the server.
	   - Header: The HTTP headers sent by the server.
	   - ContentLength: The length of the response body.
	   - Request: The original request that was sent.

	   Printing the type of the response object shows it's of type *http.Response.
	*/
	fmt.Printf("Response type= %T\n", response)

	/*
	   Always close the connection after a web request to avoid resource leaks.
	   The Body field in the response is an io.ReadCloser, which is a reader that
	   needs to be closed after reading the content to free up system resources.
	   Using defer ensures that the response body is closed when the main function exits.
	*/
	defer response.Body.Close()

	/*
	   Reading the response body:
	   The response body is read using io.ReadAll, which reads all the data from
	   the response body and returns it as a byte slice (dataBytes).
	   Handling any potential errors that may occur during the read operation.
	*/
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	/*
	   Converting the response body to a string and printing it:
	   The byte slice (dataBytes) is converted to a string using the string function.
	   This allows us to print the content of the response body in a human-readable format.
	*/
	fmt.Println(string(dataBytes))
}
