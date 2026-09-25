package main

import (
	"fmt"
	"net/http"
	"github.com/bidmybl/go-backend-practice/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Greet)
	mux.HandleFunc("/users", handler.UsersHandler)

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
