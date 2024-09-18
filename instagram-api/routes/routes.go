package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/iamaditya21/instagram-api/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/users", handlers.CreateUser)
    r.GET("/users/:id", handlers.GetUser)
    r.POST("/posts", handlers.CreatePost)
    r.GET("/posts/:id", handlers.GetPost)
    r.GET("/posts/users/:id", handlers.ListUserPosts)

    return r
}
