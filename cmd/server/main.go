package main

import (
	"fmt"
	"github.com/bidmybl/go-backend-practice/internal/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	userHandler := handler.UserHandler{}

	mux.HandleFunc("/", handler.Greet)
	mux.HandleFunc("/users", userHandler.Users)

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
