package models

type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"` // Store hashed in production
}

type Post struct {
    ID              string `json:"id"`
    Caption         string `json:"caption"`
    ImageURL        string `json:"image_url"`
    PostedTimestamp string `json:"posted_timestamp"`
    UserID          string `json:"user_id"` // Reference to the user
}
