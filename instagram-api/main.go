package main

import (
    "net/http"
    "github.com/iamaditya21/instagram-api/database"
    "github.com/iamaditya21/instagram-api/routes"
    "github.com/gorilla/mux"
)

func main() {
    database.InitDB() // Initialize the database connection
    router := mux.NewRouter()
    
    routes.UserRoutes(router)
    routes.PostRoutes(router)

    http.ListenAndServe(":8000", router)
}
