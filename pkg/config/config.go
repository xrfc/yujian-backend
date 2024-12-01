package config

import (
	"log" // use default log before we init logger

	"github.com/spf13/viper"

	"yujian-backend/pkg/model"
)

var Config model.AppConfig

// initDBConfig 初始化数据库配置。
func initDBConfig() {
	dbConfig := Config.DB
	// todo[xinhui] 从viper中获取键值对,赋值给dbConfig各个属性
	// e.g. dbConfig.Host = viper.GetString() ...
}

func InitConfig() {
	// 初始化 viper
	viper.SetConfigName("config")  // 配置文件名称（不带扩展名）
	viper.SetConfigType("yaml")    // 配置文件类型
	viper.AddConfigPath("config/") // 配置文件路径

	// 允许使用环境变量
	viper.AutomaticEnv()

	initDBConfig()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
