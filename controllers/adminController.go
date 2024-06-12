package controllers

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// code to handle login
	utils.SuccessResponse(c, "Login successful")
}

func Logout(c *gin.Context) {
	// code to handle logout
	utils.SuccessResponse(c, "Logout successful")
}

func Register(c *gin.Context) {
	// code to handle registration
	utils.SuccessResponse(c, "Registration successful")
}
