package routers

import (
	"blog/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRouter(c *gin.RouterGroup) {
	adminGroup := c.Group("/admin")
	{
		adminGroup.POST("/login", controllers.Login)
		adminGroup.POST("/register", controllers.Register)
	}
}
