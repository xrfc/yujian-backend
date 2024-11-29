package config

import (
	"log" // use default log before we init logger
	"yujian-backend/pkg/model"

	"github.com/spf13/viper"
)

var Config model.AppConfig

func initDBConfig() {

}

func initConfig() {
	// 初始化 viper
	viper.SetConfigName("config")  // 配置文件名称（不带扩展名）
	viper.SetConfigType("yaml")    // 配置文件类型
	viper.AddConfigPath("config/") // 配置文件路径

	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
