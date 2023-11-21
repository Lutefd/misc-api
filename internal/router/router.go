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
        w.Write([]byte("API Root"))
    })

    // User routes
    router.Route("/api/v1/users", func(r chi.Router) {
        r.Get("/", controllers.GetAllUsers)                   // GET /api/v1/users
        r.Post("/create", controllers.CreateUser) 
        r.Route("/{id}", func(r chi.Router) {
            r.Get("/", controllers.GetUserByID)  
            r.Put("/", controllers.UpdateUser)          
            r.Delete("/", controllers.DeleteUser)     
        })
    })

    // Album routes
    router.Route("/api/v1/albums", func(r chi.Router) {
        r.Get("/", controllers.GetAllAlbums)                 
        r.Post("/create", controllers.CreateAlbum)            
        r.Route("/{id}", func(r chi.Router) {
            r.Get("/", controllers.GetAlbumByID)             
            r.Put("/", controllers.UpdateAlbum)              
            r.Delete("/", controllers.DeleteAlbum)        
        
        })
    })

    // User Albums routes
    router.Route("/api/v1/{user_id}/albums", func(r chi.Router) {
        r.Get("/", controllers.GetAlbumsByUserID)  
        r.Delete("/{album_id}", controllers.RemoveAlbumFromUser) 
    })
    // Add the route that serves the Swagger UI

    return router
}