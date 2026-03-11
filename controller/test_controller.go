package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestController 极简测试控制器
type TestController struct{}

// GetTest 处理GET请求，返回"test"
func (tc *TestController) GetTest(c *gin.Context) {
	// 仅返回固定响应，无任何业务逻辑
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "test", // 核心：返回test
	})
}
