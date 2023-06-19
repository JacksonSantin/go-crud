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


// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Vinícius Boscardin
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
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