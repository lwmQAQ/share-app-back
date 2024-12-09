package svc

import (
	"fmt"
	"resource-server/config"
	"resource-server/internal/cache/urlcache"
	"resource-server/internal/dao"
	"resource-server/internal/mysqldb"
	"resource-server/middleware"
	"resource-server/utils"
	"resource-server/utils/etcd"
	"resource-server/utils/redisutils"
	"resource-server/utils/rpcclient"
	"resource-server/utils/urlutils"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger        *logrus.Logger
	ServerConfig  *config.ServerConfig
	EtcdUtil      *etcd.ETCDUtil
	ESClient      *utils.ESClient
	UserRpcClient *rpcclient.UserRpcClient
	MongoUtil     *utils.MongoUtil
	RedisUtil     *redisutils.RedisUtil
	UrlUtil       *urlutils.UrlUtil
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	db := mysqldb.NewMysql(&config.Mysql)
	urldao := dao.NewUrlDaoImpl(db)
	etcd := etcd.NewETCDUtil()
	es := utils.NewESlient(&config.Elasticsearch)
	userrpc := rpcclient.NewUserRpcClient(etcd)
	mongodb := utils.NewMongoUtil(logger, &config.Mongo)
	redisutil := redisutils.NewRedisUtil(&config.Redis)
	UrlCache := urlcache.NewUrlCache(redisutil, urldao)
	baseUrl := fmt.Sprintf("%s:%d/path", config.Server.Host, config.Server.Port)
	urlutil := urlutils.NewUrlUtils(baseUrl, UrlCache)
	return &ServiceContext{
		Logger:        logger,
		ServerConfig:  config,
		EtcdUtil:      etcd,
		ESClient:      es,
		UserRpcClient: userrpc,
		MongoUtil:     mongodb,
		RedisUtil:     redisutil,
		UrlUtil:       urlutil,
	}
}
