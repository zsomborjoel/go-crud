package main

import (
	"github.com/zsomborjoel/go-crud/internal/initializers"
	"github.com/zsomborjoel/go-crud/internal/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
