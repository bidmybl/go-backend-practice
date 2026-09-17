package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

var users []User

func user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid JSON")
			return
		}
		if user.Name == "" || user.Age == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Name or Age(or both) is incorrect")
			return
		}
		users = append(users, user)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := UserResponse{
			Message: "User Created",
			Data:    user,
		}
		json.NewEncoder(w).Encode(response)
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/user", user)
	http.ListenAndServe(":8080", nil)

}
