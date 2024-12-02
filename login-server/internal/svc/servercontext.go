package svc

import (
	"login-server/config"
	"login-server/internal/cache"
	"login-server/internal/dao"
	"login-server/internal/mysqldb"
	"login-server/middleware"
	"login-server/utils"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger       *logrus.Logger
	ServerConfig config.ServerConfig
	Emailutils   *utils.EmailUtils
	JWTUtil      *utils.JWTUtil
	RedisUtil    *utils.RedisUtil
	UserDao      dao.UserDao

	UserTokenCache *cache.UserTokenCache
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	mysql := mysqldb.NewMysql(&config.Mysql)
	dao.NewUserDaoImpl(mysql)

	redisutil := utils.NewRedisUtil(&config.Redis)
	email := utils.NewEmailUtils(&config.EmailConfig)
	jwt := utils.NewJWTUtil(&config.JWTConfig)

	usertokencache := cache.NewUserTokenCache(redisutil)
	return &ServiceContext{
		JWTUtil:        jwt,
		Emailutils:     email,
		Logger:         logger,
		RedisUtil:      redisutil,
		UserTokenCache: usertokencache,
	}

}
