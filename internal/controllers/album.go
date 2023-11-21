package controllers

import (
	"challenge-api/internal/helpers"
	"challenge-api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var album services.Album

func GetAllAlbums( w http.ResponseWriter, r *http.Request) {
	albums, err := album.GetAllAlbums()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all albums: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"albums": albums}, nil)
}

func GetAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	album, err := album.GetAlbumByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting album by id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"album": album}, nil)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var albumData services.Album
	err := json.NewDecoder(r.Body).Decode(&albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumCreated, err := album.CreateAlbum(albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error creating album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"album": albumCreated}, nil)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var albumData services.Album
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumUpdated, err := album.UpdateAlbum(id, albumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error updating album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"album": albumUpdated}, nil)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := album.DeleteAlbum(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Album deleted successfully"}, nil)
}

func GetAlbumsByUserID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")
	albums, err := album.GetUserAlbums(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting user albums: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"albums": albums}, nil)
}

func AddAlbumToUser(w http.ResponseWriter, r *http.Request) {
	var userAlbumData services.UserAlbum
	err := json.NewDecoder(r.Body).Decode(&userAlbumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error decoding user album: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	albumAdded, err := album.AddAlbumToUser(userAlbumData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error adding album to user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"album": albumAdded}, nil)
}

func RemoveAlbumFromUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	albumID:= chi.URLParam(r, "album_id")
	err := album.RemoveAlbumFromUser(userID, albumID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error removing album from user: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"message": "Album removed from user successfully"}, nil)
}
