package model

type User struct {
	ID   int    `json:"id"`
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
