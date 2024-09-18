package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open("postgres", "host=localhost port=5432 user=root dbname=instagram password=root sslmode=disable")
    if err != nil {
        panic("Failed to connect to database")
    }
    DB.AutoMigrate(&User{}, &Post{})
}

type User struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}

type Post struct {
    ID            uint   `json:"id" gorm:"primary_key"`
    Caption       string `json:"caption"`
    ImageURL      string `json:"image_url"`
    PostedAt      string `json:"posted_at"`
    UserID        uint   `json:"user_id"`
}
