package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/go-crud/pkg/initializers"
	"github.com/zsomborjoel/go-crud/pkg/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	v1 := r.Group("/api")
	routes.PostsRegister(v1.Group("/posts"))

	r.Run()
}
