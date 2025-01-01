package mysqldb

import (
	"fmt"
	"log"
	"time"
	"tools-back/config"

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
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 自动迁移，创建 User 表
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	// 配置连接池
	// 获取原生的 sql.DB 对象以进行连接池配置
	sqlDB, err := db.DB() // 获取底层的 *sql.DB 对象
	if err != nil {
		log.Fatalf("获取底层 sql.DB 对象失败: %v", err)
	}
	// 获取原生的 sql.DB 对象
	sqlDB.SetMaxOpenConns(10)                  // 设置最大打开连接数
	sqlDB.SetMaxIdleConns(5)                   // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 设置连接的最大生命周期
	return db
}
