package models

type Post struct {
    ID       uint   `json:"id"`
    Caption  string `json:"caption"`
    ImageURL string `json:"image_url"`
    PostedAt string `json:"posted_at"`
    UserID   uint   `json:"user_id"`
}
