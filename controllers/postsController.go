package controllers

import (
	"github.com/gin-gonic/gin"
	"example/go-crud/initializers"
	"example/go-crud/models"
)

// PostsCreate godoc
// @Summary Create a new post
// @Description Create a new post with a title and body
// @Tags posts
// @Accept json
// @Produce json
// @Param Body body string true "Post body"
// @Param Title body string true "Post title"
// @Success 200 "Post created successfully"
// @Failure 400 "Failed to create post"
// @Router /posts [post]
func PostsCreate(c *gin.Context) {
	// Get data of req body
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	c.Bind(&body)

	// Create Post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostsShow godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 "Posts finded"
// @Router /posts [get]
func PostsShow(c *gin.Context) {
	// Get thw Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// PostsIndex godoc
// @Summary Get a specific post
// @Description Get a specific post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 "Post finded"
// @Router /posts/{id} [get]
func PostsIndex(c *gin.Context) {
		// Get id off URL
		id := c.Param("id")

		// Get the Posts
		var post models.Post
		initializers.DB.Find(&post, id)

		// Return it
		c.JSON(200, gin.H{
			"post": post,
		})
}

// PostsUpdate godoc
// @Summary Update a post
// @Description Update a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param Body body string true "Post body"
// @Param Title body string true "Post title"
// @Success 200 "Post updated successfully"
// @Failure 400 "Failed to update post"
// @Router /posts/{id} [put]
func PostsUpdate(c *gin.Context) {
	// Get id off URL
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
			Title: body.Title, Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostsDelete godoc
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 "Post deleted successfully"
// @Router /posts/{id} [delete]
func PostsDelete(c *gin.Context) {
	// Get the id off URL
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}