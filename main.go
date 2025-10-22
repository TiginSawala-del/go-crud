package main

import (
	"github.com/TiginSawala-del/go-crud.git/controllers"
	"github.com/TiginSawala-del/go-crud.git/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionToDB()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostCreate)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.GET("/posts", controllers.PostIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.DELETE("/posts/:id", controllers.PostDelete)

	for _, route := range router.Routes() {
		println(route.Method, route.Path, route.Handler)
	}

	router.Run(":8001")

}
