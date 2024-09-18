package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/iamaditya21/instagram-api/instagram-api/routes" // Update with your actual path
)

func main() {
    r := mux.NewRouter()

    // Set up routes
    routes.UserRoutes(r)
    routes.PostRoutes(r)

    // Start server
    log.Fatal(http.ListenAndServe(":8000", r))
}
