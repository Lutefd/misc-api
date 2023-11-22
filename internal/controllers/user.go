package controllers

import (
	"challenge-api/internal/helpers"
	"challenge-api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type Message struct {
	Message string `json:"message"`
}

var user services.User

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieves a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} services.UsersList
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user.GetAllUsers()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all users: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"users": users}, nil)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieves a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} services.User
// @Router /users/{id} [get]
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := user.GetUserByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting user by id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"user": user}, nil)
}

// CreateUser godoc
// @Summary Create a user
// @Description Creates a new user
// @Tags users
// @Accept json
// @Produce json
// @Param userData body services.UserPayload true "User Data"
// @Success 201 {object} services.User
// @Router /users/create [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData services.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error decoding user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userCreated, err := user.CreateUser(userData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error creating user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusCreated, helpers.Envelop{"user": userCreated}, nil)
	
}

// UpdateUser godoc
// @Summary Update a user
// @Description Updates an existing user identified by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param userData body services.UserPayload true "User Data"
// @Success 200 {object} services.User
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userData services.User
	id:= chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error decoding user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userUpdated, err := user.UpdateUser(id, userData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error updating user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"user": userUpdated}, nil)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user identified by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Message
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := user.DeleteUser(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"user": Message{"User deleted"}}, nil)
}