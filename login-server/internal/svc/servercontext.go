package svc

import (
	"login-server/config"
	"login-server/middleware"
	"login-server/utils"
)

type ServiceContext struct {
	ServerConfig config.ServerConfig
	emailutils   utils.EmailUtils
}

func NewServerContext() *ServiceContext {
	logger := middleware.NewLogger()
	config := config.ReaderConfig(logger)
	email := utils.NewEmailUtils(&config.EmailConfig)
	return &ServiceContext{
		emailutils: *email,
	}

}
