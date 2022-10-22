package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zsomborjoel/go-crud/initializers"
	"github.com/zsomborjoel/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return response
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get the posts
	var post models.Post
	initializers.DB.Find(&post, id)

	// return response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// get the data off the request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// delete post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)
}
