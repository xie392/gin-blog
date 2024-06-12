package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	//var post models.Post
	//id := c.Param("id")
	//if err := database.DB.First(&user, id).Error; err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	//	return
	//}
	//c.JSON(http.StatusOK, user)
	//// TODO: Implement
	c.JSON(200, gin.H{
		"message": "Get all posts",
	})
}

func CreatePost(c *gin.Context) {
	//var post models.Post
	//if err := c.ShouldBindJSON(&post); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//if err := database.DB.Create(&post).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusCreated, post)
	// TODO: Implement
	c.JSON(201, gin.H{
		"message": "Create a new post",
	})
}
