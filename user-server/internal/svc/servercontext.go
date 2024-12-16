package svc

import (
	"user-server/config"
	"user-server/internal/cache"
	"user-server/internal/dao"
	"user-server/internal/mysqldb"
	"user-server/middleware"
	"user-server/utils"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger       *logrus.Logger
	ServerConfig *config.ServerConfig
	Emailutil    *utils.EmailUtil
	JWTUtil      *utils.JWTUtil
	EtcdUtil     *utils.ETCDUtil
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

	redis := utils.NewRedisUtil(&config.Redis)
	etcd := utils.NewETCDUtil(&config.Etcd)
	email := utils.NewEmailUtils(&config.Email)
	jwt := utils.NewJWTUtil(&config.JWT)

	usertokencache := cache.NewUserTokenCache(redis)
	codecache := cache.NewCodeCache(redis)
	userinfocache := cache.NewUserInfoCache(redis, userdao)
	return &ServiceContext{
		ServerConfig:   config,
		JWTUtil:        jwt,
		Emailutil:      email,
		Logger:         logger,
		RedisUtil:      redis,
		EtcdUtil:       etcd,
		UserTokenCache: usertokencache,
		CodeCache:      codecache,
		UserInfoCache:  userinfocache,
		UserDao:        userdao,
	}

}
