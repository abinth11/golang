package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	UserId    string   `json:"userId"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Address   *Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

var Users []User

// * middleware , helper
func (u *User) IsEmpty() bool {
	return u.Email == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/users/all", getAllUsers).Methods("GET")
	r.HandleFunc("/users/save", saveUser).Methods("POST")
	r.HandleFunc("/users/{userId}", getUserById).Methods("GET")
	r.HandleFunc("/users/update/{userId}", updateUser).Methods("PUT")
	r.HandleFunc("/users/delete/{userId}", deleteUser).Methods("DELETE")
	log.Println("Server is starting on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Page</h1>"))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	for _, user := range Users {
		if user.UserId == params["userId"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No user found with id " + params["userId"])
}

func saveUser(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid user data")
		return
	}

	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.IsEmpty() {
		json.NewEncoder(w).Encode("Provide valid user data")
		return
	}

	var userId = strconv.Itoa(rand.Intn(1000))
	user.UserId = userId
	Users = append(Users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var newUserData User
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&newUserData)
	for i, user := range Users {
		if user.UserId == params["userId"] {
			Users[i].FirstName = newUserData.FirstName
			Users[i].LastName = newUserData.LastName
			Users[i].Email = newUserData.Email
			Users[i].Address = newUserData.Address
			json.NewEncoder(w).Encode(Users[i])
			return
		}
	}
	json.NewEncoder(w).Encode("No user found with id " + params["userId"])
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	w.Header().Set("Content-Type", "application/json")
	
	for i, user := range Users {
		if user.UserId == userId {
			Users = append(Users[:i], Users[i+1:]...)
			json.NewEncoder(w).Encode(Users)
			return
		}
	}

	http.Error(w, "User not found with id "+params["userId"], http.StatusNotFound)
}
