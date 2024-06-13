package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("/api")

	PostRouter(api)
	AdminRouter(api)

	return r
}
