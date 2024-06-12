package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	PostRouter(api)
	AdminRouter(api)

	return r
}
