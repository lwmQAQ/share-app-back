package middleware

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	// 设置日志格式
	log.SetFormatter(&logrus.JSONFormatter{})

	//日志存储目录
	logDir := filepath.Join("./logs")
	// 创建日期目录（如果不存在）
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("无法创建日志目录: %v", err)
	}

	// 创建日志文件
	infoFile, err := os.OpenFile(filepath.Join(logDir, "info.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("无法创建info日志文件: %v", err)
	}
	errorFile, err := os.OpenFile(filepath.Join(logDir, "error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("无法创建error日志文件: %v", err)
	}

	// 创建文件钩子，将不同级别的日志写入不同文件
	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  io.MultiWriter(os.Stdout, infoFile),
		logrus.ErrorLevel: io.MultiWriter(os.Stdout, errorFile),
	}

	// 将日志钩子添加到日志对象
	log.AddHook(lfshook.NewHook(pathMap, &logrus.JSONFormatter{}))

	return log
}
