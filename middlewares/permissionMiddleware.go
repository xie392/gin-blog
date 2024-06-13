package middlewares

import (
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "未提供认证信息")
			c.Abort()
			return
		}

		token, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "认证信息无效")
			c.Abort()
			return
		}

		c.Set("user_id", token)

		fmt.Println("Authorization:", token)

		c.Next()
	}
}
