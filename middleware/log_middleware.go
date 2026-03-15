package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志中间件（仅打印请求信息，无复杂逻辑）
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 前置处理
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 放行请求
		c.Next()

		// 后置处理
		cost := time.Since(start)
		status := c.Writer.Status()
		fmt.Printf("[%s] %s %d %v\n", method, path, status, cost)
	}
}
