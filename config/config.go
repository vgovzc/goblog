// config/config.go
package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// 全局配置实例（供其他模块调用）
var GlobalConfig *AppConfig

// AppConfig 配置结构体（与 app.yaml 结构一一对应）
type AppConfig struct {
	Server   ServerConfig   `mapstructure:"server"` // 修复：标签拼写正确
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"` // 关键：mapstructure 不是 mapsturcture
}

// ServerConfig 服务配置
type ServerConfig struct {
	Port    string        `mapstructure:"port"`
	Mode    string        `mapstructure:"mode"`
	Timeout time.Duration `mapstructure:"timeout"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`
	Dsn             string        `mapstructure:"dsn"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level     string `mapstructure:"level"`
	Path      string `mapstructure:"path"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxBackup int    `mapstructure:"max_backup"`
}

// InitConfig 初始化 Viper 配置
func InitConfig() {
	// 1. 配置 Viper
	v := viper.New()
	// 设置配置文件路径（项目根目录的 config 文件夹）
	v.AddConfigPath("./config")
	// 设置配置文件名（不带后缀）
	v.SetConfigName("app")
	// 设置配置文件类型（yaml）
	v.SetConfigType("yaml")

	// 2. 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败：%v", err)
	}

	// 3. 将配置解析到结构体
	var appConfig AppConfig
	if err := v.Unmarshal(&appConfig); err != nil {
		log.Fatalf("解析配置文件到结构体失败：%v", err)
	}

	// 4. 赋值给全局配置
	GlobalConfig = &appConfig

	// 验证：打印配置（开发阶段用，生产可注释）
	fmt.Printf("配置加载成功：%+v\n", GlobalConfig)
}
