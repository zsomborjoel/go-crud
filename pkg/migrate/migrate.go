package main

import (
	"github.com/zsomborjoel/go-crud/pkg/initializers"
	"github.com/zsomborjoel/go-crud/pkg/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
