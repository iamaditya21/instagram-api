package handlers

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/iamaditya21/instagram-api/instagram-api/models" // Update with your actual path
)

// In-memory storage for demonstration
var users []models.User
var posts []models.Post

// Get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode("Invalid input")
        return
    }
    user.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID
    users = append(users, user)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// Get a user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, user := range users {
        if user.ID == params["id"] {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("User not found")
}

// Create a new post
func createPost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode("Invalid input")
        return
    }
    post.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID
    post.PostedTimestamp = "2023-09-18T15:04:05Z" // Use current time in production
    posts = append(posts, post)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

// Get a post by ID
func getPost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, post := range posts {
        if post.ID == params["id"] {
            json.NewEncoder(w).Encode(post)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode("Post not found")
}

// List all posts by a user
func getUserPosts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    var userPosts []models.Post
    for _, post := range posts {
        if post.UserID == params["id"] {
            userPosts = append(userPosts, post)
        }
    }
    json.NewEncoder(w).Encode(userPosts)
}
