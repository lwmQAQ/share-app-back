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
	ServerConfig *config.ServerConfig
	Emailutil    *utils.EmailUtil
	JWTUtil      *utils.JWTUtil
	RedisUtil    *utils.RedisUtil
	UserDao      dao.UserDao

	UserTokenCache *cache.UserTokenCache
	CodeCache      *cache.CodeCache
	UserInfoCache  *cache.UserInfoCache
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	mysql := mysqldb.NewMysql(&config.Mysql)
	userdao := dao.NewUserDaoImpl(mysql)

	redisutil := utils.NewRedisUtil(&config.Redis)
	email := utils.NewEmailUtils(&config.Email)
	jwt := utils.NewJWTUtil(&config.JWT)

	usertokencache := cache.NewUserTokenCache(redisutil)
	codecache := cache.NewCodeCache(redisutil)
	userinfocache := cache.NewUserInfoCache(redisutil, userdao)
	return &ServiceContext{
		ServerConfig:   config,
		JWTUtil:        jwt,
		Emailutil:      email,
		Logger:         logger,
		RedisUtil:      redisutil,
		UserTokenCache: usertokencache,
		CodeCache:      codecache,
		UserInfoCache:  userinfocache,
		UserDao:        userdao,
	}

}
