package main

import (
	"github.com/zsomborjoel/go-crud/initializers"
	"github.com/zsomborjoel/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
