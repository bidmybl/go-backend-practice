package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

var users []User

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "invalid JSON",
			})
			return
		}

		if user.Name == "" || user.Age <= 0 {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Message: "name or age is incorrect",
			})
			return
		}

		users = append(users, user)

		response := UserResponse{
			Message: "User Created",
			Data:    user,
		}

		writeJSON(w, http.StatusCreated, response)

	case http.MethodGet:
		if filterName := r.URL.Query().Get("name"); filterName != "" {
			var answer []User

			for _, filteredUser := range users {
				if filteredUser.Name == filterName {
					answer = append(answer, filteredUser)
				}
			}

			writeJSON(w, http.StatusOK, answer)
		} else {
			writeJSON(w, http.StatusOK, users)
		}

	default:
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{
			Message: "method not allowed",
		})
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", greet)
	mux.HandleFunc("/users", usersHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server started on", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
