package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bidmybl/go-backend-practice/internal/model"
	"github.com/bidmybl/go-backend-practice/internal/response"
	"net/http"
)

type UserHandler struct {
	users []model.User
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

		if user.Name == "" || user.Age <= 0 {
			response.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
				Message: "name or age is incorrect",
			})
			return
		}

		h.users = append(h.users, user)

		userResponse := model.UserResponse{
			Message: "User Created",
			Data:    user,
		}

		response.WriteJSON(w, http.StatusCreated, userResponse)

	case http.MethodGet:
		if filterName := r.URL.Query().Get("name"); filterName != "" {
			var answer []model.User

			for _, filteredUser := range h.users {
				if filteredUser.Name == filterName {
					answer = append(answer, filteredUser)
				}
			}

			response.WriteJSON(w, http.StatusOK, answer)
		} else {
			response.WriteJSON(w, http.StatusOK, h.users)
		}

	default:
		response.WriteJSON(w, http.StatusMethodNotAllowed, model.ErrorResponse{
			Message: "method not allowed",
		})
	}
}
