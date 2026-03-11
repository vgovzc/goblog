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

		// 放行请求
		c.Next()

		// 后置处理
		cost := time.Since(start)
		fmt.Printf("请求日志（空架子）：path=%s, cost=%v\n", path, cost)
	}
}
