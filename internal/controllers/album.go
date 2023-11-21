package controllers

import (
	"challenge-api/internal/helpers"
	"challenge-api/internal/services"
	"net/http"
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