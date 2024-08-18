package main

import (
	"github.com/mKepka16/go-crud/initializers"
	"github.com/mKepka16/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
