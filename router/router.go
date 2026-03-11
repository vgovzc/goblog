package router

import (
	"goblog/config"
	"goblog/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter 极简路由注册
func SetupRouter() *gin.Engine {
	// 从配置取Gin模式
	gin.SetMode(config.GlobalConfig.Server.Mode)
	r := gin.Default() // 默认包含日志+恢复中间件

	// 注册测试路由：GET /test → 调用控制器返回test
	testGroup := r.Group("/test")
	{
		tc := &controller.TestController{}
		testGroup.GET("", tc.GetTest)
	}

	return r
}
