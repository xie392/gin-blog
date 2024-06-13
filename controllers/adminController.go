package controllers

import (
	"blog/models"
	"blog/services"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.Admin
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	result, err := services.GetUser(user.Username, user.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	utils.SuccessResponse(c, result)
}

func Logout(c *gin.Context) {
	// code to handle logout
	utils.SuccessResponse(c, "Logout successful")
}

func Register(c *gin.Context) {
	var user models.Admin
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusConflict, "Username already exists")
		return
	}

	utils.SuccessResponse(c, result)
}
