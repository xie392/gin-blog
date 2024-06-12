package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	resp := Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func ErrorResponse(c *gin.Context, code int, message string) {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	c.JSON(code, resp)
}
