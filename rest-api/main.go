package main

import (
	"github.com/gin-gonic/gin"
	"rest.api/controllers"
	"rest.api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run()
}
