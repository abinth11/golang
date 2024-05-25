package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
commands
go mod tidy
go list
go list all
go get -u github.com/gorilla/mux
go mod verify
go mod init
go list -m all
go run filename.go
go build path
go list -m -versions github.com/gorilla/mux
go mod why github.com/gorilla/mux
go mod graph
go mod edit -go 1.20
go mod edit -module
go mod vendor
go run -mod=vendor main.go
*/

func main() {
	fmt.Println("Server example")
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Response from server</h1>"))
}
