package services

import (
	"github.com/mKepka16/go-crud/dtos"
	"github.com/mKepka16/go-crud/initializers"
	"github.com/mKepka16/go-crud/models"
)

func CreatePost(newPost dtos.NewPost) (dtos.Post, error) {
	post := models.Post{Title: newPost.Title, Body: newPost.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		return dtos.Post{}, result.Error
	}

	return dtos.Post{ID: post.ID, Title: post.Title, Body: post.Body, CreatedAt: post.CreatedAt}, nil
}

func GetPosts() ([]dtos.Post, error) {
	var posts []models.Post
	result := initializers.DB.Select("ID", "Title", "Body", "CreatedAt").Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	var dtosPosts []dtos.Post
	for _, post := range posts {
		dtosPosts = append(dtosPosts, dtos.Post{ID: post.ID, Title: post.Title, Body: post.Body, CreatedAt: post.CreatedAt})
	}

	return dtosPosts, nil
}

func GetPostByID(id string) (dtos.Post, error) {
	var post models.Post
	result := initializers.DB.Select("ID", "Title", "Body", "CreatedAt").Find(&post, id)
	if result.Error != nil {
		return dtos.Post{}, result.Error
	}

	return dtos.Post{ID: post.ID, Title: post.Title, Body: post.Body, CreatedAt: post.CreatedAt}, nil
}
