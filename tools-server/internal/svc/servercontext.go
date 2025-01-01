package svc

import (
	"tools-back/config"
	"tools-back/internal/dao"
	"tools-back/internal/mysqldb"
	"tools-back/internal/rpcclient"
	"tools-back/internal/types"
	"tools-back/middleware"
	"tools-back/utils"
	"tools-back/utils/etcd"

	"github.com/sirupsen/logrus"
)

type ServerContext struct {
	Logger              *logrus.Logger
	ServerConfig        *config.ServerConfig
	MinioClient         *utils.MinioClient
	RpcChan             *chan *types.RpcChanMessage // 修改为匹配类型
	FrontSSEMessageChan *chan *types.TranslationTaskResp
	EtcdUtil            *etcd.ETCDUtil
	TaskDao             dao.TaskDao
	ToolRpcClient       *rpcclient.ToolRpcClient
}

func NewServerContext(rpcchan *chan *types.RpcChanMessage) *ServerContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	minio := utils.NewMinioClient(&config.Minio)
	db := mysqldb.NewMysql(&config.Mysql)
	taskdao := dao.NewTaskDao(db)
	etcd := etcd.NewETCDUtil()
	toolrpcclient := rpcclient.NewTaskRpcClient(etcd)
	ssechan := make(chan *types.TranslationTaskResp)

	return &ServerContext{
		Logger:              logger,
		ServerConfig:        config,
		MinioClient:         minio,
		RpcChan:             rpcchan,
		TaskDao:             taskdao,
		EtcdUtil:            etcd,
		ToolRpcClient:       toolrpcclient,
		FrontSSEMessageChan: &ssechan,
	}
}
