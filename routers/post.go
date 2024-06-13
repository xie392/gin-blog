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
		postGroup.GET("", controllers.GetAllPosts)
		postGroup.GET("/:id", controllers.GetPost)
		postGroup.POST("/", permissionMiddleware, controllers.CreatePost)
		postGroup.PUT("/:id", permissionMiddleware, controllers.UpdatePost)
		postGroup.DELETE("/:id", permissionMiddleware, controllers.DeletePost)
	}
}
