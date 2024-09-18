package routes

import (
    "github.com/gorilla/mux"
    "github.com/iamaditya21/instagram-api/instagram-api/handlers" // Update with your actual path
)

func UserRoutes(router *mux.Router) {
    router.HandleFunc("/users", handlers.createUser).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.getUser).Methods("GET")
    router.HandleFunc("/users", handlers.getUsers).Methods("GET")
}

func PostRoutes(router *mux.Router) {
    router.HandleFunc("/posts", handlers.createPost).Methods("POST")
    router.HandleFunc("/posts/{id}", handlers.getPost).Methods("GET")
    router.HandleFunc("/posts/users/{id}", handlers.getUserPosts).Methods("GET")
}
