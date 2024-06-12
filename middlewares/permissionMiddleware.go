package middlewares

import "github.com/gin-gonic/gin"

func PermissionMiddleware() gin.HandlerFunc {
	// 在这里编写权限中间件的代码
	return func(c *gin.Context) {
		// 在这里进行JWT身份验证
		// 如果验证失败，则返回未授权的响应
		// 如果验证成功，则继续执行下一个处理程序
		// 示例代码：
		// if 验证失败 {
		//     c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		//     c.Abort()
		//     return
		// }
		c.Next()
	}
}
