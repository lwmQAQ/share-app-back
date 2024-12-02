package utils

import (
	"chatroom-server/jinx/jiface"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

/*全局配置类*/
type Application struct {
	/*
		Server
	*/
	Server struct {
		TcpServer jiface.IServer `yaml:"tcp-server"`
		Host      string         `yaml:"host"` //IP
		Port      int            `yaml:"port"` //端口
		Name      string         `yaml:"name"` //服务名称
		/*
			Jinx
		*/
		Version         string `yaml:"version"`            //版本
		MaxConn         int    `yaml:"max-conn"`           //最大连接数
		MaxPackageSize  uint32 `yaml:"max-package-size"`   // 单次数据包最大值
		WokerPollSize   uint32 `yaml:"worker-poll-size"`   //消息队列的数量
		MaxTaskQueueNum uint32 `yaml:"max-task-queue-num"` //每个消息队列中消息的最大值
	} `yaml:"server"`
}

/*
定义一个全局对外的
*/
var MyApplication *Application

/*
加载配置文件到全局对象
*/
func (g *Application) LoadConfig() {
	// 设置配置文件名和路径
	viper.SetConfigName("application") // 配置文件名（不含扩展名）
	viper.SetConfigType("yml")         // 配置文件类型
	viper.AddConfigPath("./config")    // 配置文件路径（当前目录）

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 解析配置文件到结构体
	if err := viper.Unmarshal(&MyApplication); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	fmt.Println("配置文件已加载:", MyApplication)
}

func init() {
	// 初始化全局对象
	MyApplication = &Application{}

	// 默认配置
	MyApplication.Server.TcpServer = nil // 可以根据实际需要设置为默认实现
	MyApplication.Server.Name = "JinxServerApp"
	MyApplication.Server.Version = "v0.4"
	MyApplication.Server.Port = 8999
	MyApplication.Server.Host = "0.0.0.0"
	MyApplication.Server.MaxConn = 1000
	MyApplication.Server.MaxPackageSize = 4096
	MyApplication.Server.WokerPollSize = 10
	MyApplication.Server.MaxTaskQueueNum = 1024
	// 从配置文件加载覆盖默认值
	MyApplication.LoadConfig()
}
