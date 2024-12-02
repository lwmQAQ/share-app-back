package mysqldb

import (
	"fmt"
	"log"
	"login-server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(dbConfig *config.MysqlConfig) *gorm.DB {

	// 使用传入的 DatabaseConfig 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	// 使用 GORM 打开数据库连接
	var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 自动迁移，创建 User 表
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	return DB
}
