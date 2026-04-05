package router

import (
	"goblog/config"
	"goblog/controller"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.GlobalConfig.Server.Mode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
	))

	// 注册测试路由：GET /test → 调用控制器返回test
	testGroup := r.Group("/test")
	{
		tc := &controller.TestController{}
		testGroup.GET("", tc.GetTest)
	}

	// 博客文章路由
	articleGroup := r.Group("/api/article")
	{
		articleGroup.GET("/list", controller.GetArticleList)
		articleGroup.GET("/:id", controller.GetArticleByID)
		articleGroup.POST("/create", controller.CreateArticle)
		articleGroup.PUT("/:id", controller.UpdateArticle)
		articleGroup.DELETE("/:id", controller.DeleteArticle)
	}

	// 用户路由
	userGroup := r.Group("/api/user")
	{
		userGroup.GET("/list", controller.GetUserList)
		userGroup.GET("/:id", controller.GetUserByID)
		userGroup.POST("/create", controller.CreateUser)
		userGroup.PUT("/:id", controller.UpdateUser)
		userGroup.DELETE("/:id", controller.DeleteUser)
	}

	// 评论路由
	commentGroup := r.Group("/api/comment")
	{
		commentGroup.GET("/list", controller.GetCommentList)
		commentGroup.GET("/:id", controller.GetCommentByID)
		commentGroup.POST("/create", controller.CreateComment)
		commentGroup.PUT("/:id", controller.UpdateComment)
		commentGroup.DELETE("/:id", controller.DeleteComment)
	}

	return r
}
