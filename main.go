package main

import (
	"goblog/config"
	"goblog/router"
	"log"
)

func main() {
	// 1. 初始化 Viper 配置（第一步必须加载）
	config.InitConfig()

	// 2. 初始化数据库（依赖 Viper 配置）
	config.InitDB()

	// 3. 初始化路由
	r := router.SetupRouter()

	// 4. 启动服务
	port := config.GlobalConfig.Server.Port
	log.Printf("极简服务启动成功：http://127.0.0.1:%s/test", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务启动失败：%v", err)
	}
}
