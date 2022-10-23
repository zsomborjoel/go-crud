package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/go-crud/pkg/handlers"
)

func PostsRegister(router *gin.RouterGroup) {
	router.POST("", handlers.PostsCreate)
	router.POST("/:id", handlers.PostsUpdate)
	router.GET("", handlers.PostsIndex)
	router.GET("/:id", handlers.PostsShow)
	router.DELETE("/:id", handlers.PostsDelete)
}
