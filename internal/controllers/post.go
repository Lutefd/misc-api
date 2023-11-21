package controllers

import (
	"challenge-api/internal/helpers"
	"challenge-api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var post services.Post

func GetAllPosts(w http.ResponseWriter, r *http.Request)  {
	posts, err := post.GetAllPosts()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting all posts: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"posts": posts}, nil)
}

func GetPostByID(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	post, err := post.GetPostByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting post by id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": post}, nil)
}

func CreatePost(w http.ResponseWriter, r *http.Request)  {
	var postData services.Post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postCreated, err := post.CreatePost(postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error creating post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": postCreated}, nil)
}

func UpdatePost(w http.ResponseWriter, r *http.Request)  {
	var postData services.Post
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error parsing post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postUpdated, err := post.UpdatePost(id, postData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error updating post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": postUpdated}, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	err := post.DeletePost(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error deleting post: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"post": "deleted"}, nil)
}

func GetPostsByUserID(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	posts, err := post.GetPostsByUserID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error getting posts by user id: ", err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelop{"posts": posts}, nil)
}