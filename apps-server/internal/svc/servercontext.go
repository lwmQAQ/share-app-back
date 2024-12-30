package svc

import (
	"apps-server/config"
	"apps-server/middleware"

	"github.com/sirupsen/logrus"
)

type ServiceContext struct {
	Logger       *logrus.Logger
	ServerConfig *config.ServerConfig
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	return &ServiceContext{
		ServerConfig: config,
	}

}
