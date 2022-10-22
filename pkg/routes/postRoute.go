package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/go-crud/pkg/controllers"
)

func PostsRegister(router *gin.RouterGroup) {
	router.POST("", controllers.PostsCreate)
	router.POST("/:id", controllers.PostsUpdate)
	router.GET("", controllers.PostsIndex)
	router.GET("/:id", controllers.PostsShow)
	router.DELETE("/:id", controllers.PostsDelete)
}
