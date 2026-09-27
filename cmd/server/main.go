package main

import (
	"fmt"
	"github.com/bidmybl/go-backend-practice/internal/handler"
	"github.com/bidmybl/go-backend-practice/internal/service"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	userService := service.NewUserService()
	userHandler := handler.UserHandler{
		Service: userService,
	}

	mux.HandleFunc("/users", userHandler.Users)
	mux.HandleFunc("/users/", userHandler.GetUser)

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
