package routers

import (
	"blog/controllers"
	"blog/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRouter(c *gin.RouterGroup) {
	permissionMiddleware := middlewares.PermissionMiddleware()

	postGroup := c.Group("/post")
	{
		postGroup.GET("/:id", controllers.GetPosts)
		postGroup.POST("/", permissionMiddleware, controllers.CreatePost)
	}
}
