package main

import (
	"fmt"
	"log"
	"net/http"

	movieRoutes "github.com/abinth11/gowithmongodb/routes"
)

func main() {
	r := movieRoutes.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("======")
	fmt.Println("Server is listening on port 4000")
}
