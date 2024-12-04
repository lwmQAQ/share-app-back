package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReaderConfig(logger *logrus.Logger) *ServerConfig {
	var config ServerConfig
	viper.SetConfigName("application") // 配置文件名，不带扩展名
	viper.SetConfigType("yml")         // 配置文件类型
	viper.AddConfigPath("./config")    // 指定配置文件所在的子目录

	// 加载配置文件内容
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("配置文件读取失败: %v", err)
	}

	// 解析配置内容
	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatalf("配置解析失败: %v", err)
	}
	fmt.Println("配置文件", config)
	return &config
}
