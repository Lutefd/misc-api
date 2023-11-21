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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user.GetAllUsers()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all users: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"users": users}, nil)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := user.GetUserByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting user by id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"user": user}, nil)
}


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
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"user": userCreated}, nil)
	
}

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