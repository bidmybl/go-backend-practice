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

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func user(w http.ResponseWriter, r *http.Request) {
	u := User{
		Name: r.URL.Query().Get("name"),
		Age:  r.URL.Query().Get("age"),
	}
	if u.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is required!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/user", user)
	http.ListenAndServe(":8080", nil)
}
