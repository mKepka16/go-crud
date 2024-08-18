package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mKepka16/go-crud/controllers"
	"github.com/mKepka16/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.CreatePost)
	r.Run() // listen and serve on 0.0.0.0:8080
}
