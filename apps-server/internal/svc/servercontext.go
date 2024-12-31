package svc

import (
	"apps-server/config"
	"apps-server/internal/dao"
	"apps-server/middleware"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger       *logrus.Logger
	ServerConfig *config.ServerConfig
	AppsDao      dao.AppsDao
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	appsdao := dao.NewAppsDaoImpl(&config.Mysql)
	return &ServiceContext{
		ServerConfig: config,
		AppsDao:      appsdao,
	}
}
