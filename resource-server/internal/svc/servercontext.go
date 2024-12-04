package svc

import (
	"resource-server/config"
	"resource-server/middleware"
	"resource-server/utils"
	"resource-server/utils/etcd"
	"resource-server/utils/rpcclient"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger        *logrus.Logger
	ServerConfig  *config.ServerConfig
	EtcdUtil      *etcd.ETCDUtil
	ESClient      *utils.ESClient
	UserRpcClient *rpcclient.UserRpcClient
	MongoUtil     *utils.MongoUtil
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	etcd := etcd.NewETCDUtil()
	es := utils.NewESlient(&config.Elasticsearch)
	userrpc := rpcclient.NewUserRpcClient(etcd)
	mongodb := utils.NewMongoUtil(logger, &config.Mongo)
	return &ServiceContext{
		Logger:        logger,
		ServerConfig:  config,
		EtcdUtil:      etcd,
		ESClient:      es,
		UserRpcClient: userrpc,
		MongoUtil:     mongodb,
	}
}
