package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/bidmybl/go-backend-practice/internal/model"
	"github.com/bidmybl/go-backend-practice/internal/response"
	"github.com/bidmybl/go-backend-practice/internal/service"
)

type UserHandler struct {
	Service *service.UserService
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
		if filterName != "" {
			users := h.Service.GetUsersByName(filterName)
			response.WriteJSON(w, http.StatusOK, users)
			return
		}
		users := h.Service.GetUsers()
		response.WriteJSON(w, http.StatusOK, users)

	default:
		response.WriteJSON(w, http.StatusMethodNotAllowed, model.ErrorResponse{
			Message: "method not allowed",
		})
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
			Message: "invalid user id",
		})
		return
	}
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			response.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
				Message: "user not found",
			})
			return
		}
	}

	response.WriteJSON(w, http.StatusOK, user)
}
