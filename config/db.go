// config/db.go
package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局 DB 实例（供其他模块调用）
var DB *gorm.DB

// InitDB 初始化数据库连接（依赖 Viper 配置）
func InitDB() {
	// 确保配置已加载
	if GlobalConfig == nil {
		log.Fatal("请先调用 InitConfig() 加载配置")
	}

	// 从全局配置获取数据库信息
	dbConfig := GlobalConfig.Database

	// 自定义 GORM 日志（方便调试）
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到控制台
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别（Info 显示所有 SQL）
			Colorful:      true,        // 彩色输出
		},
	)

	// 连接数据库（使用配置文件的 DSN）
	db, err := gorm.Open(mysql.Open(dbConfig.Dsn), &gorm.Config{
		Logger: newLogger, // 注入日志
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	// 设置连接池（使用配置文件的参数）
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败：%v", err)
	}
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	// 赋值给全局 DB
	DB = db
	log.Println("数据库连接成功！")
}
