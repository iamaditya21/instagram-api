package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/iamaditya21/instagram-api/database"
)

func CreatePost(c *gin.Context) {
    var post database.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    post.PostedAt = time.Now().Format(time.RFC3339)
    database.DB.Create(&post)
    c.JSON(http.StatusCreated, post)
}

func GetPost(c *gin.Context) {
    id := c.Param("id")
    var post database.Post
    if err := database.DB.First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    c.JSON(http.StatusOK, post)
}

func ListUserPosts(c *gin.Context) {
    userId := c.Param("id")
    var posts []database.Post
    database.DB.Where("user_id = ?", userId).Find(&posts)
    
    c.JSON(http.StatusOK, posts)
}
