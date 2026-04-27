package main

import (
	"example/go-crud/initializers"
	"example/go-crud/controllers"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
	_ "example/go-crud/docs"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


// @title GO API - CRUD Posts
// @version 1.0.0
// @contact.name Jackson Dhanyel Santin
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)

	r.GET("/posts", controllers.PostsShow)
	r.GET("/posts/:id", controllers.PostsIndex)

	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}