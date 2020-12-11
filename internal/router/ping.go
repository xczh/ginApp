package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ping-pong测试
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "pong"})
	}
}
