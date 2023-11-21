package router

import (
	"challenge-api/internal/controllers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)


func Routes() http.Handler{
    router := chi.NewRouter()
    router.Use(middleware.Recoverer)
    router.Use(middleware.Logger)
    router.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
    }))
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Api Desafio do Candidato Sensedia"))
    })
    // Users routes
    router.Get("/api/v1/users", controllers.GetAllUsers)
    router.Get("/api/v1/users/{id}", controllers.GetUserByID)
    router.Put("/api/v1/users/{id}", controllers.UpdateUser)
    router.Put("/api/v1/users/{id}", controllers.DeleteUser)
    router.Post("/api/v1/users/create", controllers.CreateUser)

    // Albums routes
    router.Get("/api/v1/albums", controllers.GetAllAlbums)


  

    // Add the route that serves the Swagger UI

    return router
}