package svc

import (
	"oss-server/config"
	"oss-server/middleware"
	"oss-server/utils"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger       *logrus.Logger
	ServerConfig *config.ServerConfig
	MinioClient  utils.MinioClient
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	minioclient := utils.NewMinioClient(&config.Minio)

	return &ServiceContext{
		ServerConfig: config,
		Logger:       logger,
		MinioClient:  *minioclient,
	}

}
