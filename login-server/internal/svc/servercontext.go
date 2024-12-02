package svc

import (
	"login-server/config"
	"login-server/internal/dao"
	"login-server/internal/mysqldb"
	"login-server/middleware"
	"login-server/utils"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger       *logrus.Logger
	ServerConfig config.ServerConfig
	Emailutils   utils.EmailUtils
	JWTUtils     utils.JWTUtil
	UserDao      dao.UserDao
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	mysql := mysqldb.NewMysql(&config.Mysql)
	dao.NewUserDaoImpl(mysql)

	email := utils.NewEmailUtils(&config.EmailConfig)
	jwt := utils.NewJWTUtil(&config.JWTConfig)
	return &ServiceContext{
		JWTUtils:   *jwt,
		Emailutils: *email,
		Logger:     logger,
	}

}
