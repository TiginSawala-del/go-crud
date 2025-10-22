package main

import (
	"github.com/TiginSawala-del/go-crud.git/initializers"
	"github.com/TiginSawala-del/go-crud.git/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectionToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}