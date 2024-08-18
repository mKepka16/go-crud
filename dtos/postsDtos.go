package dtos

import "time"

type NewPost struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type Post struct {
	ID        uint      `json:"id" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Body      string    `json:"body" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
