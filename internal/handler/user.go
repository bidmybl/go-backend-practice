package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bidmybl/go-backend-practice/internal/model"
	"github.com/bidmybl/go-backend-practice/internal/response"
	"github.com/bidmybl/go-backend-practice/internal/service"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
				Message: "invalid JSON",
			})
			return
		}

		err = h.Service.CreateUser(user)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
				Message: err.Error(),
			})
			return
		}

		userResponse := model.UserResponse{
			Message: "User Created",
			Data:    user,
		}
		response.WriteJSON(w, http.StatusCreated, userResponse)

	case http.MethodGet:
		filterName := r.URL.Query().Get("name")
		users := h.Service.GetUsers(filterName)
		response.WriteJSON(w, http.StatusOK, users)

	default:
		response.WriteJSON(w, http.StatusMethodNotAllowed, model.ErrorResponse{
			Message: "method not allowed",
		})
	}
}
