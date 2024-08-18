package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mKepka16/go-crud/dtos"
	"github.com/mKepka16/go-crud/services"
)

func CreatePost(c *gin.Context) {
	var body dtos.NewPost
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := services.CreatePost(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
	posts, err := services.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	id := c.Params.ByName("id")
	posts, err := services.GetPostByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
