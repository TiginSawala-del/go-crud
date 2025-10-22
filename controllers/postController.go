package controllers

import (
	"math"
	"strconv"

	"github.com/TiginSawala-del/go-crud.git/initializers"
	"github.com/TiginSawala-del/go-crud.git/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"error":   false,
		"message": "Berhasil Menambahkan Data",
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	var total int64

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if pageInt < 1 {
		pageInt = 1
	}

	offset := (pageInt - 1) * limitInt

	initializers.DB.Model(&models.Post{}).Count(&total)

	initializers.DB.Limit(limitInt).Offset(offset).Find(&posts)

	totalPages := int(math.Ceil(float64(total) / float64(limitInt)))

	c.JSON(200, gin.H{
		"status":      200,
		"error":       false,
		"message":     "Berhasil Mengambil Data",
		"currentPage": pageInt,
		"totalPages":  totalPages,
		"totalItems":  total,
		"data":        posts,
	})
}

func PostShow(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"status":  200,
		"error":   false,
		"message": "Berhasil Mengambil Data",
		"data":    post,
	})

}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	// body request
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"status":  200,
		"error":   false,
		"message": "Berhasil Merubah Data",
	})

}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"status":  200,
		"error":   false,
		"message": "Berhasil Menghapus Data",
	})

}
